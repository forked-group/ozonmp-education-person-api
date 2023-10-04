package repo

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type (
	eventType = education.EventType
	person    = education.Person
)

const (
	created = education.Created
	updated = education.Updated
	removed = education.Removed

	deferred  = education.Deferred
	processed = education.Processed
)

var _ interfaces.PersonRepo = (*DummyRepo)(nil)
var _ interfaces.PersonRepo = (*Repo)(nil)
