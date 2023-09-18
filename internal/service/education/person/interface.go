package person

import (
	"github.com/ozonmp/omp-bot/internal/model/education"
)

type PersonService interface {
	Describe(PersonID uint64) (*education.Person, error)
	List(cursor uint64, limit uint64) ([]education.Person, error)
	Create(education.Person) (uint64, error)
	Update(PersonID uint64, Person education.Person) error
	Remove(PersonID uint64) (bool, error)
}
