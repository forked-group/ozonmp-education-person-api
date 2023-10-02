package education

import (
	"encoding/json"
	"log"
	"time"
)

type Person struct {
	ID         uint64
	FirstName  string
	MiddleName string
	LastName   string
	Birthday   time.Time
	Sex
	Education
}

func (p Person) String() string {
	const op = "Person.String"

	buf, err := json.Marshal(p)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}

	return string(buf)
}
