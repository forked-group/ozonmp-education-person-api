package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/path"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	education_ "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type (
	person    = education_.Person
	sex       = education_.Sex
	education = education_.Education
)

var (
	parseSex       = education_.ParseSex
	parseEducation = education_.ParseEducation
	parseDate      = education_.ParseDate
)

type (
	commandPath  = path.CommandPath
	callbackPath = path.CallbackPath
)

type personService = interfaces.PersonService

var _ interfaces.PersonCommander = (*Commander)(nil)
