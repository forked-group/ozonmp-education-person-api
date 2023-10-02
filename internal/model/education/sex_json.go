package education

import (
	"bytes"
	"fmt"
	"strconv"
)

func (sx Sex) MarshalJSON() ([]byte, error) {
	const op = "Sex.MarshalJSON"

	if sx == 0 {
		return []byte(`null`), nil
	}

	if !sx.Valid() {
		return nil, fmt.Errorf("%s: unknown value %d", op, sx)
	}

	return []byte(strconv.Quote(sx.String())), nil
}

func (sx *Sex) UnmarshalJSON(text []byte) error {
	const op = "Sex.UnmarshalJSON"

	if bytes.Equal(text, []byte(`null`)) {
		// no-op rtfm
		return nil
	}

	s := string(text)

	if s == `""` || s == `0` { // is zero
		*sx = 0
		return nil
	}

	var err error

	if s[0] == '"' {
		s, err = strconv.Unquote(s)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	v, err := ParseSex(s)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	*sx = v
	return nil
}
