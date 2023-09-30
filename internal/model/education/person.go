package education

import "time"

type Person struct {
	ID         uint64
	FistName   string
	MiddleName string
	LastName   string
	Birthday   time.Time
	Sex
	Education
}

type Education int

const (
	_ Education = iota
	// Preschool Дошкольное
	Preschool
	// PrimaryGeneral Начальное общее — 1—4 классы
	PrimaryGeneral
	// BasicGeneral Основное общее — 5—9 классы
	BasicGeneral
	// SecondaryGeneral Среднее общее — 10—11 классы
	SecondaryGeneral
	// SecondaryVocational Среднее профессиональное
	SecondaryVocational
	// Higher1 Высшее I степени — бакалавриат
	Higher1
	// Higher2 Высшее II степени — специалитет, магистратура
	Higher2
	// Higher3 Высшее III степени — подготовка кадров высшей квалификации
	Higher3
)

type Sex int

const (
	_ Sex = iota
	Male
	Female
)

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
