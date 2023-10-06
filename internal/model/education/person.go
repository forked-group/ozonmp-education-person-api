package education

import (
	"strconv"
	"strings"
	"time"
)

type PersonCreate struct {
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Birthday   Date   `json:"birthday,omitempty"`
	Sex        `json:"sex,omitempty"`
	Education  `json:"education,omitempty"`
}

type Person struct {
	ID uint64 `json:"id"`
	PersonCreate
	Removed bool      `json:"removed,omitempty"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
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
