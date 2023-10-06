package repo

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type (
	person       = education.Person
	personCreate = education.PersonCreate
	personEvent  = education.PersonEvent
	eventType    = education.EventType
	eventStatus  = education.EventStatus
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
var _ interfaces.PersonEventRepo = (*EventRepo)(nil)
