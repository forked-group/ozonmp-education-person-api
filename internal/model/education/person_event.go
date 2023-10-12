package education

type EventType uint8

//go:generate stringer -type=EventType
const (
	_ EventType = iota
	Created
	Updated
	Removed
)

type EventStatus uint8

//go:generate stringer -type=EventStatus
const (
	_ EventStatus = iota
	Deferred
	Processed
)

type PersonEvent struct {
	ID     uint64      `json:"person_event_id"`
	Type   EventType   `json:"type"`
	Status EventStatus `json:"status"`
	Entity *Person     `json:"entity"`
}

type PersonEventField uint64

//go:generate stringer -type=EventStatus
const (
	PersonEventID PersonEventField = 1 << iota
	PersonEventType
	PersonEventStatus
	PersonEventEntry
)

func (mask PersonEventField) Includes(f PersonEventField) bool {
	return mask&f != 0
}
