package servers

import (
	"bufio"
	"context"
	"fmt"
	"github.com/egorgasay/gost"
	"github.com/egorgasay/itisadb-go-sdk"
	"github.com/pkg/errors"
	"itisadb/internal/constants"
	"itisadb/internal/domains"
	"itisadb/internal/models"
	"itisadb/pkg"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Servers struct {
	servers map[int32]domains.Server
	poolCh  chan struct{}
	freeID  int32
	sync.RWMutex
}

func New(local gost.Option[domains.Server]) (*Servers, error) {
	f, err := os.OpenFile("servers", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := make(map[int32]domains.Server, 10)

	if local.IsSome() {
		serv := local.Unwrap()
		s[serv.Number()] = serv
	}

	maxProc := runtime.GOMAXPROCS(0) * 100 // TODO: make it configurable

	servers := &Servers{
		servers: s,
		freeID:  2,
		poolCh:  make(chan struct{}, maxProc),
	}

	go servers.updateRAM()

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

const _updateRAMInterval = 1 * time.Second

func (s *Servers) updateRAM() {
	for {
		for _, cl := range s.servers {
			func() {
				defer s.Unlock()
				s.Lock()

				ctx := context.Background()
				ctxWithTimeout, cancel := context.WithTimeout(ctx, constants.ServerConnectTimeout)
				defer cancel()

				if r := cl.RefreshRAM(ctxWithTimeout); r.IsErr() {
					s.OnServerError(cl, r.Error())
				}
			}()

		}
		time.Sleep(_updateRAMInterval)
	}
}

func (s *Servers) OnServerError(cl domains.Server, err *gost.Error) {
	if err.BaseCode() != 0 {
		return
	}

	cl.IncTries()
	if cl.Tries() > constants.MaxServerTries {
		s.Disconnect(cl.Number())
	}
}

func (s *Servers) GetServer() (domains.Server, bool) {
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

	ctx := context.Background()

	var cl *itisadb.Client

	switch r := itisadb.New(ctx, address); r.Switch() { /* TODO: add creds? */
	case gost.IsOk:
		cl = r.Unwrap()
	case gost.IsErr:
		return 0, r.Error()
	}

	// add test connection

	var stClient = &RemoteServer{
		sdk: cl,
		ram: models.RAM{Available: available, Total: total},
		mu:  &sync.RWMutex{},
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
	f, err := os.OpenFile("balancer", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return 0, errors.Wrapf(ErrInternal, "can't open file: balancer, %v", err.Error())
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

	var servers = make([]string, len(s.servers))
	for _, cl := range s.servers {
		r := cl.RAM()
		servers = append(servers, fmt.Sprintf("s#%d Avaliable: %d MB, Total: %d MB", cl.Number(), r.Available, r.Total))
	}

	return servers
}

func (s *Servers) DeepSearch(ctx context.Context, key string, opts models.GetOptions) (models.Value, error) {
	s.RLock()
	defer s.RUnlock()

	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var out = make(chan models.Value)
	var once sync.Once

	var finished = func(v models.Value) {
		once.Do(func() {
			defer close(out)
			out <- v
		})
	}

	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	for _, cl := range s.servers {
		c := cl

		go pkg.WithContext(ctx, func() error {
			defer wg.Done()

			r := c.GetOne(ctxCancel, key, opts)
			if r.IsErr() {
				s.OnServerError(c, r.Error())
				return nil
			}

			finished(r.Unwrap())
			return nil
		}, s.poolCh, nil)
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
		return models.Value{}, constants.ErrNotFound
	}
}

func (s *Servers) GetServerByID(number int32) (domains.Server, bool) {
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

	wg.Add(len(s.servers))

	set := func(server domains.Server, number int32) {
		defer wg.Done()
		if r := server.SetOne(ctx, key, val, opts); r.IsErr() {
			s.OnServerError(server, r.Error())

			mu.Lock()
			failedServers = append(failedServers, number)
			mu.Unlock()
			return
		}

		server.ResetTries()
	}

	for n, serv := range s.servers {
		s.poolCh <- struct{}{}
		go func(serv domains.Server, n int32) {
			set(serv, n)
			<-s.poolCh
		}(serv, n)
	}
	wg.Wait()

	return failedServers
}

func (s *Servers) DelFromAll(ctx context.Context, key string, opts models.DeleteOptions) (atLeastOnce bool) {
	var wg sync.WaitGroup

	wg.Add(len(s.servers))

	del := func(server domains.Server, number int32) {
		defer wg.Done()
		if r := server.DelOne(ctx, key, opts); r.IsErr() {
			s.OnServerError(server, r.Error())
			return
		}
		atLeastOnce = true

		server.ResetTries()
	}

	for n, serv := range s.servers {
		s.poolCh <- struct{}{}
		go func(serv domains.Server, n int32) {
			del(serv, n)
			<-s.poolCh
		}(serv, n)
	}
	wg.Wait()

	return atLeastOnce
}
