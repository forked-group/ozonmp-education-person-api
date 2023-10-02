package education

import (
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
