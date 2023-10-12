package repo

import (
	"context"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	"github.com/rs/zerolog/log"
)

type EventRepo interface {
	Lock(n uint64) ([]model.PersonEvent, error)
	Unlock(eventIDs []uint64) error

	Add(events []model.PersonEvent) error
	Remove(eventIDs []uint64) error
}

var _ EventRepo = EventRepoAdapter{}

type EventRepoAdapter struct {
	Repo *repo.PersonEventRepo
}

func (r EventRepoAdapter) Lock(n uint64) ([]model.PersonEvent, error) {
	const op = "EventRepoAdapter.Lock"

	ctx := context.TODO()
	batch, err := r.Repo.Lock(ctx, n)
	if err != nil {
		log.Error().Err(err).Str("op", op).Msgf("fail")
	}
	return batch, err
}

func (r EventRepoAdapter) Unlock(eventIDs []uint64) error {
	const op = "EventRepoAdapter.Unlock"

	ctx := context.TODO()
	_, err := r.Repo.Unlock(ctx, eventIDs)
	if err != nil {
		log.Error().Err(err).Str("op", op).Msgf("fail")
	}
	return err
}

func (r EventRepoAdapter) Add(events []model.PersonEvent) error {
	//TODO implement me
	panic("implement me")
}

func (r EventRepoAdapter) Remove(eventIDs []uint64) error {
	const op = "EventRepoAdapter.Remove"

	ctx := context.TODO()
	_, err := r.Repo.Remove(ctx, eventIDs)
	if err != nil {
		log.Error().Err(err).Str("op", op).Msgf("fail")
	}
	return err
}
