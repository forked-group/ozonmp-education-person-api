package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type (
	person      = model.Person
	personField = model.PersonField
	personRepo  = interfaces.PersonRepo
)

var _ interfaces.PersonService = (*Service)(nil)
