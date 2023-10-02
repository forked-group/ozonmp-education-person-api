package education

import (
	"bytes"
	"encoding/json"
)

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
	if bytes.Equal(text, []byte(`null`)) {
		// no-op rtfm
		return nil
	}

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
