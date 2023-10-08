package education

import "fmt"

type EventType uint8

//go:generate stringer -type=EventType
const (
	_ EventType = iota
	Created
	Updated
	Removed
)

func ParseEventType(s string) (EventType, error) {
	const op = "ParseEventStatus"
	switch s {
	case Created.String():
		return Created, nil
	case Updated.String():
		return Updated, nil
	case Removed.String():
		return Removed, nil
	}
	return 0, fmt.Errorf("%s: unknown event status: %q", op, s)
}

type EventStatus uint8

//go:generate stringer -type=EventStatus
const (
	_ EventStatus = iota
	Deferred
	Processed
)

func ParseEventStatus(s string) (EventStatus, error) {
	const op = "ParseEventStatus"
	switch s {
	case Deferred.String():
		return Deferred, nil
	case Processed.String():
		return Processed, nil
	}
	return 0, fmt.Errorf("%s: unknown event status: %q", op, s)
}

type PersonEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Person
}

type PersonEventField uint64

const (
	PersonEventID PersonEventField = 1 << iota
	PersonEventType
	PersonEventStatus
	PersonEventEntry
)
