package repo

import (
	"context"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/service/education/person"
)

var _ Repo = (*dummyRepo)(nil)

type dummyRepo struct {
	personService person.PersonService
}

func NewDummyRepo(personService person.PersonService) Repo {
	return dummyRepo{
		personService: personService,
	}
}

func (d dummyRepo) DescribePerson(ctx context.Context, personID uint64) (*education.Person, error) {
	return d.personService.Describe(personID)
}

func (d dummyRepo) ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]education.Person, error) {
	return d.personService.List(cursor, limit)
}

func (d dummyRepo) CreatePerson(ctx context.Context, person education.Person) (uint64, error) {
	return d.personService.Create(person)
}

func (d dummyRepo) UpdatePerson(ctx context.Context, personID uint64, person education.Person) error {
	return d.personService.Update(personID, person)
}

func (d dummyRepo) RemovePerson(ctx context.Context, personID uint64) (bool, error) {
	return d.personService.Remove(personID)
}
