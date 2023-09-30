package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

// Repo is DAO for Person
type Repo interface {
	DescribePerson(ctx context.Context, personID uint64) (*model.Person, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribePerson(ctx context.Context, personID uint64) (*model.Person, error) {
	return nil, nil
}
