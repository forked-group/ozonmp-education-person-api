package person

type Person struct {
	ID uint64
}

type EventType uint8

const (
	Created EventType = iota
	Updated
	Removed
)

type EventStatus uint8

const (
	Deferred EventStatus = iota
	Processed
)

type PersonEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Person
}
