package sender

import "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"

type EventSender interface {
	Send(person *education.PersonEvent) error
}
