package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type (
	person       = education.Person
	personCreate = education.PersonCreate
	personRepo   = interfaces.PersonRepo
)

var _ interfaces.PersonService = (*Service)(nil)
