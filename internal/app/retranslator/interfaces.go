package retranslator

//go:generate mockgen -destination=./mocks/event_locker.go . EventLocker
type EventLocker interface {
	Lock(n uint64) ([]event, error)
}

type EventUnlocker interface {
	Unlock(eventIDs []uint64) error
}

type EventRemover interface {
	Remove(eventIDs []uint64) error
}

//go:generate mockgen -destination=./mocks/event_sender.go . EventSender
type EventSender interface {
	Send(event *event) error
}

//go:generate mockgen -destination=./mocks/event_repo.go . EventRepo
type EventRepo interface {
	EventLocker
	EventUnlocker
	EventRemover
}
