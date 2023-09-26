package retranslator_test

import (
	"container/heap"
	"errors"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/repo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/sender"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"math/rand"
	"sync"
	"time"
)

var _ repo.EventRepo = (*dummyRepo)(nil)
var _ sender.EventSender = (*dummySender)(nil)

var errUnknownError = errors.New("unknown error")

func randTimeoutAndError(timeout time.Duration, errPer100K int) error {
	if timeout > 0 {
		// TODO: заменить на нормальное распределение
		timeout = time.Duration(rand.Int63n(int64(timeout))) + timeout/2 // timeout * (1±1/2)
		time.Sleep(timeout)
	}

	if errPer100K > 0 && rand.Intn(100_000) < errPer100K {
		return errUnknownError
	}

	return nil
}

type DummySenderCfg struct {
	Size       int
	Timeout    time.Duration
	ErrPer100K int
}

type dummySender struct {
	cfg    *DummySenderCfg
	mu     sync.Mutex
	Events []uint64
}

func (cfg *DummySenderCfg) New() *dummySender {
	events := make([]uint64, 0, cfg.Size)
	return &dummySender{
		cfg:    cfg,
		Events: events,
	}
}

func (s *dummySender) Send(events *person.PersonEvent) error {
	if err := randTimeoutAndError(s.cfg.Timeout, s.cfg.ErrPer100K); err != nil {
		return err
	}
	return s.send(events)
}

func (s *dummySender) send(event *person.PersonEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Events = append(s.Events, event.ID)
	return nil
}

type DummyRepoCfg struct {
	Size       int
	Timeout    time.Duration
	ErrPer100K int
}

type dummyRepo struct {
	cfg       *DummyRepoCfg
	mu        sync.Mutex
	Deferred  Uint64Heap
	processed []bool
	Removed   []uint64
}

func (cfg *DummyRepoCfg) New() *dummyRepo {
	n := cfg.Size

	deferred := make(Uint64Heap, n)
	for i := 0; i < n; i++ {
		deferred[i] = uint64(i + 1)
	}
	heap.Init(&deferred)

	processed := make([]bool, n+1)
	removed := make([]uint64, 0, n)

	return &dummyRepo{
		cfg:       cfg,
		Deferred:  deferred,
		processed: processed,
		Removed:   removed,
	}
}

func (s *dummyRepo) Lock(n uint64) ([]person.PersonEvent, error) {
	if err := randTimeoutAndError(s.cfg.Timeout, s.cfg.ErrPer100K); err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if n > uint64(len(s.Deferred)) {
		n = uint64(len(s.Deferred))
	}

	events := make([]person.PersonEvent, 0, n)

	for ; n > 0 && len(s.Deferred) != 0; n-- {
		// pop from deferred
		id := heap.Pop(&s.Deferred).(uint64)

		// mark as processed
		s.processed[id] = true

		// push to Events
		events = append(events, person.PersonEvent{
			ID:     id,
			Type:   person.Created,
			Status: person.Processed,
		})
	}

	return events, nil
}

func (s *dummyRepo) Unlock(eventIDs []uint64) error {
	if err := randTimeoutAndError(s.cfg.Timeout, s.cfg.ErrPer100K); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(eventIDs) - 1; i >= 0; i-- {
		id := eventIDs[i]
		if 0 <= id && id < uint64(len(s.processed)) && s.processed[id] {
			s.processed[id] = false
			heap.Push(&s.Deferred, id)
		}
	}

	return nil
}

func (s *dummyRepo) Add(events []person.PersonEvent) error {
	//TODO implement me
	panic("implement me")
}

func (s *dummyRepo) Remove(eventIDs []uint64) error {
	if err := randTimeoutAndError(s.cfg.Timeout, s.cfg.ErrPer100K); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 0; i < len(eventIDs); i++ {
		id := eventIDs[i]
		if 0 <= id && id < uint64(len(s.processed)) && s.processed[id] {
			s.processed[id] = false
			s.Removed = append(s.Removed, id)
		}
	}

	return nil
}

func (s *dummyRepo) ProcessedLen() int {
	return cap(s.Deferred) - len(s.Deferred) - len(s.Removed)
}

func (s *dummyRepo) Processed() []uint64 {
	result := make([]uint64, 0, s.ProcessedLen())

	for i, v := range s.processed {
		if v {
			result = append(result, uint64(i))
		}
	}

	return result
}

//=============================================================================
// https://pkg.go.dev/container/heap@go1.21.1#example-package-IntHeap
//

// An Uint64Heap is a min-heap of ints.
type Uint64Heap []uint64

func (h Uint64Heap) Len() int           { return len(h) }
func (h Uint64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Uint64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Uint64Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(uint64))
}

func (h *Uint64Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//=============================================================================
