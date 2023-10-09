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
	return p.repo.Describe(context.TODO(), personID)
}

func (p Service) List(cursor uint64, limit uint64) ([]person, error) {
	return p.repo.List(context.TODO(), cursor, limit)
}

func (p Service) Create(person person) (uint64, error) {
	return p.repo.Create(context.TODO(), person)
}

func (p Service) Update(personID uint64, person person, fields personField) (bool, error) {
	return p.repo.Update(context.TODO(), personID, person, fields)
}

func (p Service) Remove(personID uint64) (bool, error) {
	return p.repo.Remove(context.TODO(), personID)
}
