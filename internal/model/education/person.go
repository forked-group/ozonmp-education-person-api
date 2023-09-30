package education

import (
	"encoding/json"
	"log"
)

type Person struct {
	ID         uint64 `json:"id"`
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Birthday   *Date  `json:"birthday,omitempty"`
	Sex        `json:"sex,omitempty"`
	Education  `json:"education,omitempty"`
}

func (p Person) String() string {
	const op = "Person.String"

	buf, err := json.Marshal(p)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}
	return string(buf)
}
