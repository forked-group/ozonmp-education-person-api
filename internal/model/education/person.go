package education

import (
	"strconv"
	"strings"
	"time"
)

type Person struct {
	ID         uint64
	FirstName  string
	MiddleName string
	LastName   string
	Birthday   time.Time
	Sex
	Education
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
		sb.WriteByte(' ')
		sb.WriteString(p.Birthday.Format(DateLayout))
	}

	if p.Sex != 0 {
		sb.WriteByte(' ')
		sb.WriteString(p.Sex.String())
	}

	if p.Education != 0 {
		sb.WriteByte(' ')
		sb.WriteString(p.Education.String())
	}

	return sb.String()
}
