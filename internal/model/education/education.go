package education

import (
	"fmt"
	"strconv"
	"strings"
)

type Education int

const (
	_ Education = iota
	// Preschool Дошкольное
	Preschool
	// PrimaryGeneral Начальное общее — 1—4 классы
	PrimaryGeneral
	// BasicGeneral Основное общее — 5—9 классы
	BasicGeneral
	// SecondaryGeneral Среднее общее — 10—11 классы
	SecondaryGeneral
	// SecondaryVocational Среднее профессиональное
	SecondaryVocational
	// Higher1 Высшее I степени — бакалавриат
	Higher1
	// Higher2 Высшее II степени — специалитет, магистратура
	Higher2
	// Higher3 Высшее III степени — подготовка кадров высшей квалификации
	Higher3
)

func ParseEducation(s string) (Education, error) {
	const op = "ParseEducation"

	if n, err := strconv.Atoi(s); err == nil {
		if e := Education(n); Preschool <= e && e <= Higher3 {
			return e, nil
		}

		return 0, fmt.Errorf("%s: unknown value: %s", op, s)
	}

	for e := Preschool; e <= Higher3; e++ {
		if strings.EqualFold(s, e.String()) {
			return e, nil
		}
	}

	return 0, fmt.Errorf("%s: unknown value: %q", op, s)
}

func (e Education) String() string {
	switch e {
	case Preschool:
		return "Preschool"
	case PrimaryGeneral:
		return "PrimaryGeneral"
	case BasicGeneral:
		return "BasicGeneral"
	case SecondaryGeneral:
		return "SecondaryGeneral"
	case SecondaryVocational:
		return "SecondaryVocational"
	case Higher1:
		return "Higher1"
	case Higher2:
		return "Higher2"
	case Higher3:
		return "Higher3"
	default:
		return ""
	}
}

func (e Education) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(e.String())), nil
}
