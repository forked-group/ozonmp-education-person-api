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

func NewDate(t time.Time) *Date {
	return &Date{
		time.Date(t.Year(), t.Month(), t.Day(),
			0, 0, 0, 0, time.UTC),
	}
}

func ParseDate(s string) (Date, error) {
	if t, err := time.Parse(DateLayout, s); err != nil {
		return Date{}, err
	} else {
		return Date{t}, nil
	}
}

func (d Date) String() string {
	return d.Time.Format(DateLayout)
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(d.String())), nil
}

func (d *Date) UnmarshalJSON(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		return nil
	}

	s, err := strconv.Unquote(string(text))
	if err != nil {
		return err
	}

	v, err := ParseDate(s)
	if err != nil {
		return err
	}

	*d = v
	return nil
}
