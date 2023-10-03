package retranslator

import (
	"container/heap"
	"errors"
	"math/rand"
	"sync"
	"time"
)

var _ EventRepo = (*dummyRepo)(nil)
var _ EventSender = (*dummySender)(nil)

var ErrUnknownError = errors.New("unknown error")

func randTimeoutAndError(timeout time.Duration, errPer100K int) error {
	if timeout > 0 {
		// TODO: заменить на нормальное распределение
		timeout = time.Duration(rand.Int63n(int64(timeout))) + timeout/2 // timeout * (1±1/2)
		time.Sleep(timeout)
	}

	if errPer100K > 0 && rand.Intn(100_000) < errPer100K {
		return ErrUnknownError
	}

	return nil
}

type DummySenderCfg struct {
	Size       int
	Latency    time.Duration
	ErrPer100K int
}

type dummySender struct {
	cfg  *DummySenderCfg
	mu   sync.Mutex
	sent []uint64
}

func (cfg *DummySenderCfg) New() *dummySender {
	events := make([]uint64, 0, cfg.Size)
	return &dummySender{
		cfg:  cfg,
		sent: events,
	}
}

func (s *dummySender) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.sent)
}

func (s *dummySender) Send(events *event) error {
	if err := randTimeoutAndError(s.cfg.Latency, s.cfg.ErrPer100K); err != nil {
		return err
	}
	return s.send(events)
}

func (s *dummySender) send(event *event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sent = append(s.sent, event.ID)
	return nil
}

type DummyRepoCfg struct {
	Size       int
	Latency    time.Duration
	ErrPer100K int
}

type dummyRepo struct {
	cfg       *DummyRepoCfg
	mu        sync.Mutex
	deferred  Uint64Heap
	processed *Uint64SparseMap
	removed   []uint64
	missed    int
}

func (cfg *DummyRepoCfg) New() *dummyRepo {
	n := cfg.Size

	deferred := make(Uint64Heap, n)
	for i := 0; i < n; i++ {
		deferred[i] = uint64(i + 1)
	}
	heap.Init(&deferred)

	processed := NewUin64SpareMap(n + 1)
	removed := make([]uint64, 0, n)

	return &dummyRepo{
		cfg:       cfg,
		deferred:  deferred,
		processed: processed,
		removed:   removed,
	}
}

func (r *dummyRepo) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.deferred)
}

func (r *dummyRepo) Lock(n uint64) ([]event, error) {
	if err := randTimeoutAndError(r.cfg.Latency, r.cfg.ErrPer100K); err != nil {
		return nil, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if n > uint64(len(r.deferred)) {
		n = uint64(len(r.deferred))
	}

	events := make([]event, 0, n)

	for ; n > 0 && len(r.deferred) != 0; n-- {
		id := heap.Pop(&r.deferred).(uint64)
		r.processed.Add(id)

		events = append(events, event{
			ID:     id,
			Type:   eventTypeCreated,
			Status: eventStatusProcessed,
		})
	}

	return events, nil
}

func (r *dummyRepo) Unlock(eventIDs []uint64) error {
	if err := randTimeoutAndError(r.cfg.Latency, r.cfg.ErrPer100K); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, id := range eventIDs {
		if !r.processed.Includes(id) {
			r.missed++
		} else {
			r.processed.Delete(id)
			heap.Push(&r.deferred, id)
		}
	}

	return nil
}

func (r *dummyRepo) Add(events []event) error {
	//TODO implement me
	panic("implement me")
}

func (r *dummyRepo) Remove(eventIDs []uint64) error {
	if err := randTimeoutAndError(r.cfg.Latency, r.cfg.ErrPer100K); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < len(eventIDs); i++ {
		id := eventIDs[i]
		if !r.processed.Includes(id) {
			r.missed++
		} else {
			r.processed.Delete(id)
			r.removed = append(r.removed, id)
		}
	}

	return nil
}

// Uint64SparseMap sparse map из статьи Расса Кокса [https://research.swtch.com/sparse]
type Uint64SparseMap struct {
	dense  []uint64
	sparse []uint64
}

func NewUin64SpareMap(n int) *Uint64SparseMap {
	spare := make([]uint64, n)
	return &Uint64SparseMap{sparse: spare}
}

func (sm *Uint64SparseMap) Len() int {
	return len(sm.dense)
}

func (sm *Uint64SparseMap) Values() []uint64 {
	return sm.dense
}

func (sm *Uint64SparseMap) Add(v uint64) {
	n := len(sm.dense)
	sm.sparse[v] = uint64(n)
	sm.dense = append(sm.dense, v)
}

func (sm *Uint64SparseMap) Delete(v uint64) {
	if i := sm.sparse[v]; i < uint64(len(sm.dense)) && sm.dense[i] == v {
		n := len(sm.dense) - 1
		last := sm.dense[n]

		sm.dense[i] = last
		sm.sparse[last] = i
		sm.dense = sm.dense[:n]
	}
}

func (sm *Uint64SparseMap) Includes(v uint64) bool {
	i := sm.sparse[v]
	return i < uint64(len(sm.dense)) && sm.dense[i] == v
}

// Uint64Heap is a min-heap of ints [https://pkg.go.dev/container/heap@go1.21.1#example-package-IntHeap]
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
