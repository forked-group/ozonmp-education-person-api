package sender

import "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/person"

type EventSender interface {
	Send(person *person.PersonEvent) error
}
