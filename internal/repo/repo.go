package repo

import (
	"context"

	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

// Repo is DAO for Person
type Repo interface {
	DescribePerson(ctx context.Context, personID uint64) (*education.Person, error)
	ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]education.Person, error)
	CreatePerson(ctx context.Context, person education.Person) (uint64, error)
	UpdatePerson(ctx context.Context, personID uint64, person education.Person) error
	RemovePerson(ctx context.Context, personID uint64) (bool, error)
}

//type repo struct {
//	db        *sqlx.DB
//	batchSize uint
//}
//
//// NewRepo returns Repo interface
//func NewRepo(db *sqlx.DB, batchSize uint) Repo {
//	return &repo{db: db, batchSize: batchSize}
//}
//
//func (r *repo) DescribePerson(ctx context.Context, personID uint64) (*education.Person, error) {
//	return nil, nil
//}
