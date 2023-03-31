package usecase

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"sync"

	repo "grpc-storage/internal/memory-balancer/storage"
	"grpc-storage/pkg/api/storage"

	"github.com/tomakado/containers/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var ErrNoData = errors.New("the value is not found")

type UseCase struct {
	clients map[uint64]storage.StorageClient
	sync.RWMutex
	logger  *zap.Logger
	storage *repo.Storage
	queue   *queue.Queue[uint64]
}

func New(repository *repo.Storage, logger *zap.Logger) *UseCase {
	clients := make(map[uint64]storage.StorageClient, 10)
	return &UseCase{
		clients: clients,
		storage: repository,
		logger:  logger,
		queue:   &queue.Queue[uint64]{},
	}
}

func (uc *UseCase) Set(key string, val string) (uint64, error) {
	uc.RLock()
	defer uc.RUnlock()
	if len(uc.clients) == 0 {
		err := uc.storage.Set(key, val)
		if err != nil {
			return 0, fmt.Errorf("error while setting new pair to dbstorage with no active grpc-storages: %w", err)
		}
		return 0, nil
	}
	serverNumber := uint64(len(key)%len(uc.clients) + 1)
	cl, ok := uc.clients[serverNumber]
	if !ok || cl == nil {
		err := uc.storage.Set(key, val)
		if err != nil {
			return 0, fmt.Errorf("error while adding new pair to dbstorage with offline grpc-storage: %w", err)
		}
		return 0, nil
	}

	_, err := cl.Set(context.Background(), &storage.SetRequest{Key: key, Value: val})
	if err != nil {
		return 0, nil
	}

	return serverNumber, nil
}

func (uc *UseCase) Get(key string) (string, error) {
	uc.RLock()
	defer uc.RUnlock()

	if len(uc.clients) == 0 {
		value, err := uc.storage.Get(key)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return "", fmt.Errorf("error while getting new pair to dbstorage with no active grpc-storages: %w", ErrNoData)
			}
			return value, fmt.Errorf("error while getting new pair to dbstorage with no active grpc-storages: %w", err)
		}
		return value, nil
	}

	cl, ok := uc.clients[uint64(len(key)%(len(uc.clients))+1)]
	if !ok || cl == nil {
		value, err := uc.storage.Get(key)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return "", fmt.Errorf("error while getting new pair to dbstorage with offline grpc-storages: %w", ErrNoData)
			}
			return value, fmt.Errorf("error while getting new pair to dbstorage with offline grpc-storage: %w", err)
		}
		return value, nil
	}

	res, err := cl.Get(context.Background(), &storage.GetRequest{Key: key})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return "", err
		}
		if st.Code().String() == codes.NotFound.String() {
			return "", ErrNoData
		}

		if st.Code().String() == codes.Unavailable.String() { // connection error
			get, err := uc.storage.Get(key)
			if errors.Is(err, mongo.ErrNoDocuments) {
				return "", ErrNoData
			}
			return get, err
		}

		return "", fmt.Errorf("can't get the value from server: %w", err)
	}

	return res.Value, nil
}

func (uc *UseCase) Connect(address string) (uint64, error) {
	uc.Lock()
	defer uc.Unlock()

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}

	cl := storage.NewStorageClient(conn)

	numForReuse, ok := uc.queue.Dequeue()
	number := uint64(len(uc.clients) + 1)
	if ok {
		number = numForReuse
	}

	uc.clients[number] = cl
	return number, nil
}

func (uc *UseCase) Disconnect(number uint64) error {
	uc.RLock()
	defer uc.RUnlock()
	uc.clients[number] = nil
	uc.queue.Enqueue(number)

	return nil
}
