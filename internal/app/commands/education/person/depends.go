package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/path"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type (
	person      = model.Person
	personField = model.PersonField
)

const (
	personFirstName  = model.PersonFirstName
	personMiddleName = model.PersonMiddleName
	personLastName   = model.PersonLastName
	personBirthday   = model.PersonBirthday
	personSex        = model.PersonSex
	personEducation  = model.PersonEducation
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
