package education

import (
	"fmt"
	"strconv"
	"strings"
)

type Education int

//go:generate stringer -type=Education
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

	MinEducation = Preschool
	MaxEducation = Higher3
)

func (ed Education) Valid() bool {
	return MinEducation <= ed && ed <= MaxEducation
}

func ParseEducation(s string) (Education, error) {
	const op = "ParseEducation"

	if i, err := strconv.Atoi(s); err == nil {
		if ed := Education(i); ed.Valid() {
			return ed, nil
		}
		return 0, fmt.Errorf("%s: unknown value: %s", op, s)
	}

	for ed := MinEducation; ed <= MaxEducation; ed++ {
		if strings.EqualFold(s, ed.String()) {
			return ed, nil
		}
	}

	return 0, fmt.Errorf("%s: unknown value: %q", op, s)
}
