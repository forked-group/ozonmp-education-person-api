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

func (sx Sex) Valid() bool {
	return MinSex <= sx && sx <= MaxSex
}

func ParseSex(s string) (Sex, error) {
	const op = "ParseSex"

	if i, err := strconv.Atoi(s); err == nil {
		if sex := Sex(i); sex.Valid() {
			return sex, nil
		}
		return 0, fmt.Errorf("%s: unknown value: %s", op, s)
	}

	for sex := MinSex; sex <= MaxSex; sex++ {
		if strings.EqualFold(s, sex.String()) {
			return sex, nil
		}
	}

	return 0, fmt.Errorf("%s: unknown value: %q", op, s)
}
