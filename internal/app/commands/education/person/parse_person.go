package person

import (
	"errors"
	"fmt"
	"strings"
)

func parsePersonNames(args []string, p *personCreate) ([]string, error) {

	names := make([]string, 0, 3)
	for i := 0; i < 3 && i < len(args); i++ {
		if strings.IndexByte(args[i], '=') != -1 {
			break
		}
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

func parsePersonFields(args []string, p *personCreate) (err error) {

	for _, arg := range args {
		pos := strings.IndexByte(arg, '=')
		if pos == -1 {
			return fmt.Errorf("%q: '=' expected", arg)
		}

		name := arg[:pos]
		val := arg[pos+1:]

		switch name {
		case "first_name":
			p.FirstName = val

		case "middle_name":
			p.MiddleName = val

		case "last_name":
			p.LastName = val

		case "birthday":
			var v date
			if val == "" {
				p.Birthday = date{}
			} else if v, err = parseDate(val); err == nil {
				p.Birthday = v
			}

		case "sex":
			var v sex
			if val == "" {
				p.Sex = 0
			} else if v, err = parseSex(val); err == nil {
				p.Sex = v
			}

		case "education":
			var v education
			if val == "" {
				p.Education = 0
			} else if v, err = parseEducation(val); err == nil {
				p.Education = v
			}

		default:
			err = errors.New("unknown field")
		}

		if err != nil {
			return fmt.Errorf("%s: %v", name, err)
		}
	}

	return nil
}
