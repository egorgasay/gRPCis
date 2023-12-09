package servers

import (
	"bufio"
	"context"
	"fmt"
	"github.com/egorgasay/gost"
	"github.com/egorgasay/itisadb-go-sdk"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"itisadb/internal/constants"
	"itisadb/internal/models"
	"itisadb/pkg/api"
	"os"
	"runtime"
	"strconv"
	"sync"
)

type Servers struct {
	servers map[int32]Server
	poolCh  chan struct{}
	freeID  int32
	sync.RWMutex
}

type Server interface {
	RAM() models.RAM
	SetRAM(ram models.RAM)

	Number() int32
	Tries() uint32
	IncTries()
	ResetTries()

	Find(ctx context.Context, key string, out chan<- string, once *sync.Once, opts models.GetOptions)

	appLogic
}

type appLogic interface {
	GetOne(ctx context.Context, key string, opts ...itisadb.GetOptions) (res gost.Result[string])
	DelOne(ctx context.Context, key string, opts ...itisadb.DeleteOptions) gost.Result[bool]
	SetOne(ctx context.Context, key string, val string, opts ...itisadb.SetOptions) (res gost.Result[int32])
}

//Set(ctx context.Context, key string, value string, opts models.SetOptions) error
//Get(ctx context.Context, key string, opts models.GetOptions) (*api.GetResponse, error)
//ObjectToJSON(ctx context.Context, name string, opts models.ObjectToJSONOptions) (*api.ObjectToJSONResponse, error)
//GetFromObject(ctx context.Context, name string, key string, opts models.GetFromObjectOptions) (*api.GetFromObjectResponse, error)
//SetToObject(ctx context.Context, name string, key string, value string, opts models.SetToObjectOptions) error
//NewObject(ctx context.Context, name string, opts models.ObjectOptions) error
//Size(ctx context.Context, name string, opts models.SizeOptions) (*api.ObjectSizeResponse, error)
//DeleteObject(ctx context.Context, name string, opts models.DeleteObjectOptions) error
//Delete(ctx context.Context, Key string, opts models.DeleteOptions) error
//AttachToObject(ctx context.Context, dst string, src string, opts models.AttachToObjectOptions) error
//DeleteAttr(ctx context.Context, attr string, object string, opts models.DeleteAttrOptions) error

func New() (*Servers, error) {
	f, err := os.OpenFile("servers", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := make(map[int32]Server, 10)

	maxProc := runtime.GOMAXPROCS(0) * 100

	servers := &Servers{
		servers: s,
		freeID:  1,
		poolCh:  make(chan struct{}, maxProc),
	}

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		_, err = f.WriteString("1")
		return servers, err
	}

	line := scanner.Text()
	atoi, err := strconv.Atoi(line)
	if err != nil {
		return nil, fmt.Errorf("can't get the last used number: %w", err)
	}

	servers.freeID = int32(atoi)

	return servers, nil
}

func (s *Servers) GetServer() (Server, bool) {
	s.RLock()
	defer s.RUnlock()

	best := 0.0
	var serverNumber int32 = 0

	for num, cl := range s.servers {
		r := cl.RAM()
		if val := float64(r.Available) / float64(r.Total) * 100; val > best {
			serverNumber = num
			best = val
		}
	}

	cl, ok := s.servers[serverNumber]
	return cl, ok
}

func (s *Servers) Len() int32 {
	s.RLock()
	defer s.RUnlock()
	return int32(len(s.servers))
}

var ErrInternal = errors.New("internal error")

func (s *Servers) AddServer(address string, available, total uint64, server int32) (int32, error) {
	s.Lock()
	defer s.Unlock()

	if server == 0 {
		server = s.freeID
		s.freeID++
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, errors.Wrap(ErrInternal, err.Error())
	}

	cl := api.NewItisaDBClient(conn)

	// add test connection

	var stClient = &RemoteServer{
		client: cl,
		ram:    models.RAM{Available: available, Total: total},
		mu:     &sync.RWMutex{},
	}

	if server != 0 {
		stClient.number = server
		if server > s.freeID {
			s.freeID = server + 1
		}
	} else {
		stClient.number = s.freeID
		s.freeID++
	}

	// saving last id
	f, err := os.OpenFile("servers", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return 0, errors.Wrapf(ErrInternal, "can't open file: servers, %v", err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%d\n", s.freeID))
	if err != nil {
		return 0, errors.Wrapf(ErrInternal, "can't save last id: %v", err.Error())
	}

	s.servers[stClient.number] = stClient

	return stClient.number, nil
}

func (s *Servers) Disconnect(number int32) {
	s.Lock()
	defer s.Unlock()
	delete(s.servers, number)
}

func (s *Servers) GetServers() []string {
	s.RLock()
	defer s.RUnlock()

	var servers = make([]string, 0, 5)
	for _, cl := range s.servers {
		r := cl.RAM()
		servers = append(servers, fmt.Sprintf("s#%d Avaliable: %d MB, Total: %d MB",
			cl.Number(), r.Available, r.Total))
	}

	return servers
}

func (s *Servers) DeepSearch(ctx context.Context, key string, opts models.GetOptions) (string, error) {
	s.RLock()
	defer s.RUnlock()

	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var out = make(chan string, 1)
	defer close(out)

	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	var once sync.Once
	for _, cl := range s.servers {
		c := cl
		s.poolCh <- struct{}{}
		go func() {
			c.Find(ctxCancel, key, out, &once, opts)
			<-s.poolCh
			wg.Done()
		}()
	}

	allIsDone := make(chan struct{})

	go func() {
		wg.Wait()
		close(allIsDone)
	}()

	select {
	case v := <-out:
		cancel()
		return v, nil
	case <-allIsDone:
		return "", constants.ErrNotFound
	}
}

func (s *Servers) GetServerByID(number int32) (Server, bool) {
	s.RLock()
	defer s.RUnlock()
	srv, ok := s.servers[number]
	return srv, ok
}

func (s *Servers) Exists(number int32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.servers[number]
	return ok
}

func (s *Servers) SetToAll(ctx context.Context, key, val string, opts models.SetOptions) []int32 {
	var failedServers = make([]int32, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup

	s.RLock()
	defer s.RUnlock()

	wg.Add(len(s.servers))

	set := func(server Server, number int32) {
		defer wg.Done()
		err := server.SetOne(ctx, key, val, opts).Error()
		if err != nil {
			if server.Tries() > 2 {
				delete(s.servers, number)
			}
			server.IncTries()
			mu.Lock()
			failedServers = append(failedServers, number)
			mu.Unlock()
			return
		}

		server.ResetTries()
	}

	for n, serv := range s.servers {
		s.poolCh <- struct{}{}
		go func(serv Server, n int32) {
			set(serv, n)
			<-s.poolCh
		}(serv, n)
	}
	wg.Wait()

	return failedServers
}

func (s *Servers) DelFromAll(ctx context.Context, key string, opts models.DeleteOptions) (atLeastOnce bool) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	s.RLock()
	defer s.RUnlock()

	wg.Add(len(s.servers))

	del := func(server Server, number int32) {
		defer wg.Done()
		err := server.DelOne(ctx, key, opts).Error()
		if err != nil {
			if server.Tries() > 2 {
				delete(s.servers, number)
			}
			server.IncTries()
			mu.Lock()
			mu.Unlock()
			return
		}
		atLeastOnce = true

		server.ResetTries()
	}

	for n, serv := range s.servers {
		s.poolCh <- struct{}{}
		go func(serv Server, n int32) {
			del(serv, n)
			<-s.poolCh
		}(serv, n)
	}
	wg.Wait()

	return atLeastOnce
}
