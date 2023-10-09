package education

import (
	"fmt"
	"strconv"
	"strings"
)

type Sex int

//go:generate stringer -type=Sex
const (
	_ Sex = iota
	Female
	Male
	MinSex = Female
	MaxSex = Male
)

func (i Sex) Valid() bool {
	return MinSex <= i && i <= MaxSex
}

func ParseSex(s string) (Sex, error) {
	const op = "ParseSex"

	if i, err := strconv.Atoi(s); err == nil {
		if Sex(i).Valid() {
			return Sex(i), nil
		}
		return 0, fmt.Errorf("%s: unknown value: %s", op, s)
	}

	for i := MinSex; i <= MaxSex; i++ {
		if strings.EqualFold(s, i.String()) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("%s: unknown value: %q", op, s)
}
