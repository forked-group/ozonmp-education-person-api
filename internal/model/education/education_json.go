package education

import (
	"bytes"
	"fmt"
	"strconv"
)

func (ed Education) MarshalJSON() ([]byte, error) {
	const op = "Education.MarshalJSON"

	if ed == 0 {
		return []byte(`null`), nil
	}

	if !ed.Valid() {
		return nil, fmt.Errorf("%s: unknown value %d", op, ed)
	}

	return []byte(strconv.Quote(ed.String())), nil
}

func (ed *Education) UnmarshalJSON(text []byte) error {
	const op = "Education.UnmarshalJSON"

	if bytes.Equal(text, []byte(`null`)) {
		// no-op rtfm
		return nil
	}

	s := string(text)

	if s == `""` || s == `0` { // is zero
		*ed = 0
		return nil
	}

	var err error

	if s[0] == '"' {
		s, err = strconv.Unquote(s)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	v, err := ParseEducation(s)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	*ed = v
	return nil
}
