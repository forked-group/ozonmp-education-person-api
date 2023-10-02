package education

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func (d Date) MarshalJSON() (text []byte, err error) {
	if time.Time(d).IsZero() {
		return []byte(`null`), nil // TODO: how to tell to json that it is null?
	}
	return []byte(strconv.Quote(d.String())), nil
}

func (d *Date) UnmarshalJSON(text []byte) error {
	const op = "Date.UnmarshalJSON"

	if bytes.Equal(text, []byte(`null`)) {
		// no-op rtfm
		return nil
	}

	if bytes.Equal(text, []byte(`""`)) {
		*d = Date{}
		return nil
	}

	s, err := strconv.Unquote(string(text))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	v, err := ParseDate(s)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	*d = v
	return nil
}
