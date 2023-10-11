package person

import (
	"errors"
	"fmt"
	"strings"
)

func parsePersonNames(args []string, p *person) ([]string, error) {
	n := min(3, len(args))

	names := make([]string, 0, n)
	for i := 0; i < n && strings.IndexByte(args[i], '=') == -1; i++ {
		names = append(names, args[i])
	}

	switch len(names) {
	case 1:
		p.LastName = names[0]
		args = args[1:]

	case 2:
		p.FirstName = names[0]
		p.LastName = names[1]
		args = args[2:]

	case 3:
		p.FirstName = names[0]
		p.MiddleName = names[1]
		p.LastName = names[2]
		args = args[3:]
	}

	return args, nil
}

func parsePersonFields(args []string, p *person, f *personField) error {

	for _, arg := range args {
		pos := strings.IndexByte(arg, '=')
		if pos == -1 {
			return fmt.Errorf("%q: '=' expected", arg)
		}

		name := arg[:pos]
		val := arg[pos+1:]

		var err error

		switch name {
		case "first_name":
			p.FirstName = val
			*f |= personFirstName

		case "middle_name":
			p.MiddleName = val
			*f |= personMiddleName

		case "last_name":
			p.LastName = val
			*f |= personLastName

		case "birthday":
			if val == "" {
				p.Birthday = nil
			} else {
				p.Birthday, err = parseDate(val)
			}
			*f |= personBirthday

		case "sex":
			if val == "" {
				p.Sex = 0
			} else {
				p.Sex, err = parseSex(val)
			}
			*f |= personSex

		case "education":
			if val == "" {
				p.Education = 0
			} else {
				p.Education, err = parseEducation(val)
			}
			*f |= personEducation

		default:
			err = errors.New("unknown field")
		}

		if err != nil {
			return fmt.Errorf("%s: %v", name, err)
		}
	}

	return nil
}
