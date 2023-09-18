package person

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"math/rand"
	"sync"
	"time"
)

var testData = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
}

type DummyPersonService struct {
	data    []education.Person
	index   map[uint64]int
	deleted map[uint64]struct{}
	mu      sync.RWMutex
	rand    *rand.Rand
}

func NewDummyPersonService() *DummyPersonService {
	return &DummyPersonService{
		index:   map[uint64]int{},
		deleted: map[uint64]struct{}{},
		rand:    rand.New(rand.NewSource(time.Now().UnixMicro())),
	}
}

func NewDummyPersonServiceWithTestData() *DummyPersonService {
	s := NewDummyPersonService()

	for _, name := range testData {
		s.Create(education.Person{Name: name})
	}

	return s
}

var ErrNotFound = errors.New("not found")

func (s *DummyPersonService) Describe(personID uint64) (*education.Person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	index, ok := s.index[personID]
	if !ok {
		return nil, ErrNotFound
	}

	if _, ok := s.deleted[personID]; ok {
		return nil, ErrNotFound
	}

	// Returns a pointer to a copy of the element for security.
	// Imho, this function should return a value, but not a pointer
	person := s.data[index]
	return &person, nil
}

func (s *DummyPersonService) List(cursor uint64, limit uint64) ([]education.Person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	index := -1

	if cursor != 0 {
		var ok bool
		if index, ok = s.index[cursor]; !ok {
			return nil, ErrNotFound
		}
	}

	result := make([]education.Person, 0, limit)

	for i := index + 1; i < len(s.data) && uint64(len(result)) < limit; i++ {
		person := s.data[i]
		if _, ok := s.deleted[person.ID]; !ok {
			result = append(result, person)
		}
	}

	return result, nil
}

func (s *DummyPersonService) randID() uint64 {
	for {
		id := s.rand.Uint64()
		if _, ok := s.index[id]; !ok {
			return id
		}
	}
}

func (s *DummyPersonService) Create(person education.Person) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	person.ID = s.randID()

	s.data = append(s.data, person)
	s.index[person.ID] = len(s.data) - 1

	return person.ID, nil
}

func (s *DummyPersonService) Update(personID uint64, person education.Person) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	index, ok := s.index[personID]
	if !ok {
		return ErrNotFound
	}

	if _, ok := s.deleted[personID]; ok {
		return ErrNotFound
	}

	person.ID = personID
	s.data[index] = person

	return nil
}

func (s *DummyPersonService) Remove(personID uint64) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.index[personID]; !ok {
		return false, nil
	}

	if _, ok := s.deleted[personID]; ok {
		return false, nil
	}

	s.deleted[personID] = struct{}{}

	return true, nil
}
