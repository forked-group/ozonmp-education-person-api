package person

import (
	"context"
)

type Service struct {
	repo personRepo
}

func NewService(r personRepo) *Service {
	return &Service{
		repo: r,
	}
}

func (p Service) Describe(personID uint64) (*person, error) {
	return p.repo.DescribePerson(context.TODO(), personID)
}

func (p Service) List(cursor uint64, limit uint64) ([]person, error) {
	return p.repo.ListPerson(context.TODO(), cursor, limit)
}

func (p Service) Create(person person) (uint64, error) {
	return p.repo.CreatePerson(context.TODO(), person)
}

func (p Service) Update(personID uint64, person person) (bool, error) {
	return p.repo.UpdatePerson(context.TODO(), personID, person)
}

func (p Service) Remove(personID uint64) (bool, error) {
	return p.repo.RemovePerson(context.TODO(), personID)
}
