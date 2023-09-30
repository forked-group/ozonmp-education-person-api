package sender

type EventSender interface {
	Send(person *education.PersonEvent) error
}
