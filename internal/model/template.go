package model

// Person - template entity.
type Person struct {
	ID  uint64 `db:"id"`
	Foo uint64 `db:"foo"`
}
