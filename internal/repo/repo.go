package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/aaa2ppp/ozonmp-education-kw-person-api/internal/model"
)

// Repo is DAO for Template
type Repo interface {
	DescribeTemplate(ctx context.Context, personID uint64) (*model.Template, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeTemplate(ctx context.Context, personID uint64) (*model.Template, error) {
	return nil, nil
}
