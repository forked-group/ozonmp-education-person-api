package repo

import (
	"github.com/aaa2ppp/ozonmp-education-kw-person-api/internal/model/person"
)

type EventRepo interface {
	Lock(n uint64) ([]person.PersonEvent, error)
	Unlock(eventIDs []uint64) error

	Add(events []person.PersonEvent) error
	Remove(eventIDs []uint64) error
}
