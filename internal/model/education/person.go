package education

type Person struct {
	ID   uint64
	Name string
}

func (p Person) String() string {
	return p.Name
}
