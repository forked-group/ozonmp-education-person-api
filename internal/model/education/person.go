package education

import (
	"strconv"
	"strings"
	"time"
)

//go:generate stringer -type=PersonField
type PersonField uint64

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

func (set PersonField) Includes(f PersonField) bool {
	return set&f != 0
}

type Person struct {
	ID         uint64    `json:"person_id"`
	FirstName  string    `json:"first_name,omitempty"`
	MiddleName string    `json:"middle_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Birthday   *Date     `json:"birthday,omitempty"`
	Sex        Sex       `json:"sex,omitempty"`
	Education  Education `json:"education,omitempty"`
	Removed    bool      `json:"removed,omitempty"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
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

	if p.Birthday != nil {
		sb.WriteString(",  ")
		sb.WriteString(p.Birthday.String())
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
