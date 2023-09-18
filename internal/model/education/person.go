package education

import (
	"encoding/json"
)

type Person struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (p Person) String() string {
	buf, _ := json.Marshal(p)
	return string(buf)
}
