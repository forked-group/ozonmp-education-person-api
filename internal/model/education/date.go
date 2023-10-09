package education

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

const DateLayout = "2006-01-02"

var _ json.Marshaler = Date{}
var _ json.Unmarshaler = (*Date)(nil)

type Date struct {
	time.Time
}

func NewDate(year int, month time.Month, day int) *Date {
	return &Date{
		time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

func ParseDate(s string) (*Date, error) {
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return nil, err
	}
	return &Date{Time: t}, err
}

func (d *Date) NullTime() *time.Time {
	if d == nil {
		return nil
	}
	return &d.Time
}

func (d Date) String() string {
	return d.Format(DateLayout)
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(d.String())), nil
}

func (d *Date) UnmarshalJSON(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		return nil // RTFM
	}

	s, err := strconv.Unquote(string(text))
	if err != nil {
		return err
	}

	v, err := ParseDate(s)
	if err != nil {
		return err
	}

	*d = *v
	return nil
}
