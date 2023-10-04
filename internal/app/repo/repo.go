package repo

import "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"

type EventRepo interface {
	Lock(n uint64) ([]education.PersonEvent, error)
	Unlock(eventIDs []uint64) error

	Add(events []education.PersonEvent) error
	Remove(eventIDs []uint64) error
}
