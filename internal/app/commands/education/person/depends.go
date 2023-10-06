package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/path"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type (
	person      = model.Person
	personField = model.PersonField
	sex         = model.Sex
	education   = model.Education
	date        = model.Date
)

var (
	parseSex       = model.ParseSex
	parseEducation = model.ParseEducation
	parseDate      = model.ParseDate
)

type (
	commandPath  = path.CommandPath
	callbackPath = path.CallbackPath
)

type (
	personService = interfaces.PersonService
	commanderCfg  = interfaces.CommanderCfg
)

var _ interfaces.PersonCommander = (*Commander)(nil)
