package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type PersonService interface {
	Describe(personID uint64) (*education.Person, error)
	List(cursor uint64, limit uint64) ([]education.Person, error)
	Create(education.Person) (uint64, error)
	Update(personID uint64, person education.Person) error
	Remove(personID uint64) (bool, error)
}
