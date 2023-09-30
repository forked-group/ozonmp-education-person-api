package education

import (
	"fmt"
	"strconv"
	"strings"
)

type Sex int

const (
	_ Sex = iota
	Female
	Male
)

func PaseSex(s string) (Sex, error) {
	const op = "ParseSex"

	if n, err := strconv.Atoi(s); err == nil {
		switch sex := Sex(n); sex {
		case Female, Male:
			return sex, nil
		default:
			return 0, fmt.Errorf("%s: unknown value: %s", op, s)
		}
	}

	switch {
	case strings.EqualFold(s, Female.String()):
		return Female, nil
	case strings.EqualFold(s, Male.String()):
		return Male, nil
	default:
		return 0, fmt.Errorf("%s: unknown value: %q", op, s)
	}
}

func (s Sex) String() string {
	switch s {
	case Female:
		return "Female"
	case Male:
		return "Male"
	default:
		return ""
	}
}

func (s Sex) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(s.String())), nil
}
