package person

import (
	"errors"
	"fmt"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"strings"
)

const DateLayout = "2006-01-02"

func parsePersonNames(args []string, p *person) ([]string, error) {

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

func parsePersonFields(args []string, p *person, f *personField) (err error) {

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
			*f |= model.PersonFirstName

		case "middle_name":
			p.MiddleName = val
			*f |= model.PersonMiddleName

		case "last_name":
			p.LastName = val
			*f |= model.PersonLastName

		case "birthday":
			if val == "" {
				p.Birthday = nil
			} else {
				p.Birthday, err = model.ParseDate(val)
			}
			*f |= model.PersonBirthday

		case "sex":
			if val == "" {
				p.Sex = 0
			} else {
				p.Sex, err = parseSex(val)
			}
			*f |= model.PersonSex

		case "education":
			if val == "" {
				p.Education = 0
			} else {
				p.Education, err = parseEducation(val)
			}
			*f |= model.PersonEducation

		default:
			err = errors.New("unknown field")
		}

		if err != nil {
			return fmt.Errorf("%s: %v", name, err)
		}
	}

	return nil
}
