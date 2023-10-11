package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
)

var _ interfaces.PersonService = (*Service)(nil)
