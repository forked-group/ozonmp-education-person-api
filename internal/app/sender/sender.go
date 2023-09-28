package sender

import "github.com/aaa2ppp/ozonmp-education-kw-person-api/internal/model/person"

type EventSender interface {
	Send(person *person.PersonEvent) error
}
