package education

import (
	"strconv"
	"strings"
	"time"
)

//go:generate stringer -type=PersonField
type PersonField int

const (
	PersonID PersonField = 1 << iota
	PersonFirstName
	PersonMiddleName
	PersonLastName
	PersonBirthday
	PersonSex
	PersonEducation
	PersonRemoved
	PersonCreated
	PersonUpdated
)

func (set PersonField) IsSet(f PersonField) bool {
	return set&f != 0
}

type Person struct {
	ID         uint64
	FirstName  string
	MiddleName string
	LastName   string
	Birthday   time.Time
	Sex
	Education
	Removed bool
	Created time.Time
	Updated time.Time
}

func (p Person) String() string {
	var sb strings.Builder

	sb.WriteString(strconv.FormatUint(p.ID, 10))
	sb.WriteByte(':')

	if p.FirstName != "" {
		sb.WriteByte(' ')
		sb.WriteString(p.FirstName)
	}

	if p.MiddleName != "" {
		sb.WriteByte(' ')
		sb.WriteString(p.MiddleName)
	}

	if p.LastName != "" {
		sb.WriteByte(' ')
		sb.WriteString(p.LastName)
	}

	if !p.Birthday.IsZero() {
		sb.WriteString(",  ")
		sb.WriteString(p.Birthday.Format(DateLayout))
	}

	if p.Sex != 0 {
		sb.WriteString(", ")
		sb.WriteString(p.Sex.String())
	}

	if p.Education != 0 {
		sb.WriteString(", ")
		sb.WriteString(p.Education.String())
	}

	return sb.String()
}
