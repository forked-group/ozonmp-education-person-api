package repo

import (
	"context"
	"os"
	"sync"
	"time"
)

const envWithTestData = "WITH_TEST_DATA"

// repo will be filled with this data if envWithTestData is set
var personTestData = []person{
	{LastName: "One"},
	{LastName: "Two"},
	{LastName: "Three"},
	{LastName: "Four"},
	{LastName: "Five"},
	{LastName: "Six"},
	{LastName: "Seven"},
	{LastName: "Eight"},
	{LastName: "Nine"},
	{LastName: "Ten"},
}

type DummyRepo struct {
	batchSize uint
	data      []person
	mu        sync.RWMutex
}

func NewDummyRepo(batchSize uint) *DummyRepo {
	s := &DummyRepo{
		batchSize: batchSize,
		data:      []person{{Removed: true}}, // stub for unused index 0
	}
	if _, ok := os.LookupEnv(envWithTestData); ok {
		s.fillTestData()
	}
	return s
}

func (s *DummyRepo) fillTestData() {
	for _, person := range personTestData {
		_, _ = s.CreatePerson(context.Background(), person)
	}
}

func (s *DummyRepo) inRange(index uint64) bool {
	return 1 <= index && index < uint64(len(s.data))
}

func (s *DummyRepo) DescribePerson(ctx context.Context, personID uint64) (*person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	index := personID

	if !s.inRange(index) || s.data[index].Removed {
		return nil, nil
	}

	// we return a pointer to a copy of the entry for security
	person := s.data[index]

	return &person, nil
}

func (s *DummyRepo) ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if 0 <= limit || limit > uint64(s.batchSize) {
		limit = uint64(s.batchSize)
	}

	result := make([]person, 0, limit)

	for i := cursor; i < uint64(len(s.data)) && uint64(len(result)) < limit; i++ {
		if s.data[i].Removed {
			continue
		}
		result = append(result, s.data[i])
	}

	return result, nil
}

func (s *DummyRepo) CreatePerson(ctx context.Context, person person) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	person.ID = uint64(len(s.data))
	person.Created = time.Now()
	person.Updated = time.Now()
	s.data = append(s.data, person)

	return person.ID, nil
}

func (s *DummyRepo) UpdatePerson(ctx context.Context, personID uint64, person person) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := personID
	if !s.inRange(index) || s.data[index].Removed {
		return false, nil
	}

	person.ID = personID
	person.Updated = time.Now()
	s.data[index] = person

	return true, nil
}

func (s *DummyRepo) RemovePerson(ctx context.Context, personID uint64) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := personID
	if !s.inRange(index) || s.data[index].Removed {
		return false, nil
	}

	s.data[index].Removed = true
	s.data[index].Updated = time.Now()

	return true, nil
}
