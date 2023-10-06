package repo

import (
	"context"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"os"
	"sync"
	"time"
)

const envWithTestData = "WITH_TEST_DATA"

// repo will be filled with this data if envWithTestData is set
var personTestData = []model.Person{
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

type PersonDummyRepo struct {
	batchSize uint
	data      []model.Person
	mu        sync.RWMutex
}

func NewDummyRepo(batchSize uint) *PersonDummyRepo {
	s := &PersonDummyRepo{
		batchSize: batchSize,
		data:      []model.Person{{Removed: true}}, // stub for unused index 0
	}
	if _, ok := os.LookupEnv(envWithTestData); ok {
		s.fillTestData()
	}
	return s
}

func (s *PersonDummyRepo) fillTestData() {
	for _, person := range personTestData {
		_, _ = s.CreatePerson(context.Background(), person)
	}
}

func (s *PersonDummyRepo) inRange(index uint64) bool {
	return 1 <= index && index < uint64(len(s.data))
}

func (s *PersonDummyRepo) DescribePerson(_ context.Context, personID uint64) (*model.Person, error) {
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

func (s *PersonDummyRepo) ListPerson(_ context.Context, cursor uint64, limit uint64) ([]model.Person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if 0 <= limit || limit > uint64(s.batchSize) {
		limit = uint64(s.batchSize)
	}

	result := make([]model.Person, 0, limit)

	for i := cursor; i < uint64(len(s.data)) && uint64(len(result)) < limit; i++ {
		if s.data[i].Removed {
			continue
		}
		result = append(result, s.data[i])
	}

	return result, nil
}

func (s *PersonDummyRepo) CreatePerson(_ context.Context, p model.Person) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.ID = uint64(len(s.data))
	s.data = append(s.data, p)

	return p.ID, nil
}

func (s *PersonDummyRepo) UpdatePerson(_ context.Context, personID uint64, p model.Person, _ model.PersonField) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := personID
	if !s.inRange(index) || s.data[index].Removed {
		return false, nil
	}

	p.ID = personID
	s.data[index] = p

	return true, nil
}

func (s *PersonDummyRepo) RemovePerson(ctx context.Context, personID uint64) (bool, error) {
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
