package person

import (
	"github.com/ozonmp/omp-bot/internal/model/education"
	"github.com/ozonmp/omp-bot/internal/service"
	"sync"
)

var testData = []education.Person{
	{LastName: "Zero"},
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

var ErrNotFound = service.ErrNotFound

type dataItem struct {
	education.Person
	deleted bool
}

type DummyPersonService struct {
	data []dataItem
	mu   sync.RWMutex
}

func NewDummyPersonService() *DummyPersonService {
	return &DummyPersonService{}
}

func NewDummyPersonServiceWithTestData() *DummyPersonService {
	s := NewDummyPersonService()

	for _, person := range testData {
		s.Create(person)
	}

	return s
}

func (s *DummyPersonService) Describe(personID uint64) (*education.Person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	index := personID

	if !(0 <= index && index < uint64(len(s.data))) || s.data[index].deleted {
		return nil, ErrNotFound
	}

	// Returns a pointer to a copy of the element for security.
	// Imho, this function should return a value, but not a pointer
	person := s.data[index].Person
	return &person, nil
}

func (s *DummyPersonService) List(cursor uint64, limit uint64) ([]education.Person, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]education.Person, 0, limit)

	for i := cursor; i < uint64(len(s.data)) && uint64(len(result)) < limit; i++ {
		if s.data[i].deleted {
			continue
		}
		result = append(result, s.data[i].Person)
	}

	if len(result) == 0 {
		return nil, ErrNotFound
	}

	return result, nil
}

func (s *DummyPersonService) Create(person education.Person) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	person.ID = uint64(len(s.data))
	s.data = append(s.data, dataItem{Person: person})

	return person.ID, nil
}

func (s *DummyPersonService) Update(personID uint64, person education.Person) error {
	const op = "DummyPersonService.Update"

	s.mu.Lock()
	defer s.mu.Unlock()

	index := personID
	if s.data[index].deleted {
		return ErrNotFound
	}

	person.ID = personID
	s.data[index].Person = person

	return nil
}

func (s *DummyPersonService) Remove(personID uint64) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := personID
	if !(0 <= index && index < uint64(len(s.data))) || s.data[index].deleted {
		return false, nil
	}

	s.data[index].deleted = true

	return true, nil
}
