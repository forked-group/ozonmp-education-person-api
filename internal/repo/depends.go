package repo

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
)

var _ interfaces.PersonRepo = (*PersonDummyRepo)(nil)
var _ interfaces.PersonRepo = (*PersonRepo)(nil)
var _ interfaces.PersonEventRepo = (*PersonEventRepo)(nil)
