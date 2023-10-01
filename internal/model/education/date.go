package education

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type Date time.Time

const DateLayout = "2006-01-02"

func NewDate(t time.Time) *Date {
	if t.IsZero() {
		return nil
	}
	d := Date(t)
	return &d
}

func (d *Date) Time() time.Time {
	if d == nil {
		return time.Time{}
	}
	return time.Time(*d)
}

func ParseDate(s string) (Date, error) {
	t, err := time.Parse(DateLayout, s)
	return Date(t), err
}

func (d Date) String() string {
	return time.Time(d).Format(DateLayout)
}

func (d Date) MarshalJSON() (text []byte, err error) {
	if time.Time(d).IsZero() {
		return []byte("null"), nil // TODO: how to tell it's null?
	}
	return []byte(strconv.Quote(d.String())), nil
}

func (d *Date) UnmarshalJSON(text []byte) error {
	const op = "Date.UnmarshalJSON"

	if bytes.Equal(text, []byte("null")) {
		// no-op rtfm
		return nil
	}

	s, err := strconv.Unquote(string(text))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	*d, err = ParseDate(s)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
