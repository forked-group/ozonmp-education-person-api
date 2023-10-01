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

type jsonPerson struct {
	ID         uint64 `json:"id"`
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Birthday   *Date  `json:"birthday,omitempty"`
	Sex        `json:"sex,omitempty"`
	Education  `json:"education,omitempty"`
}

func newJsonPerson(p Person) *jsonPerson {
	return &jsonPerson{
		ID:         p.ID,
		FirstName:  p.FirstName,
		MiddleName: p.MiddleName,
		LastName:   p.LastName,
		Birthday:   NewDate(p.Birthday),
		Sex:        p.Sex,
		Education:  p.Education,
	}
}

func (p Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(newJsonPerson(p))
}

func (p *Person) UnmarshalJSON(text []byte) error {
	jp := newJsonPerson(*p)

	if err := json.Unmarshal(text, jp); err != nil {
		return err
	}

	*p = Person{
		ID:         jp.ID,
		FirstName:  jp.FirstName,
		MiddleName: jp.MiddleName,
		LastName:   jp.LastName,
		Birthday:   jp.Birthday.Time(),
		Sex:        jp.Sex,
		Education:  jp.Education,
	}

	return nil
}

func (p Person) String() string {
	const op = "Person.String"

	buf, err := json.Marshal(p)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}

	return string(buf)
}
