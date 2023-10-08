package repo

import (
	"math"
	"strconv"
	"strings"
)

func CommaList(elems []string) string {
	return strings.Join(elems, ",")
}

func ValueListTemplate(firstPlaceholderNum int, n int) string {
	elems := make([]string, n)

	for i, num := 0, firstPlaceholderNum; i < n; i, num = i+1, num+1 {
		elems[i] = "$" + strconv.Itoa(num)
	}

	return CommaList(elems)
}

func FieldListTemplate(firstPlaceholderNum int, names []string) string {
	n := len(names)
	elems := make([]string, n)

	for i, num := 0, firstPlaceholderNum; i < n; i, num = i+1, num+1 {
		elems[i] = names[i] + "=$" + strconv.Itoa(num)
	}

	return CommaList(elems)
}

const AllFields = math.MaxUint64 // usage: h := NewListsBuilder(CustomType(AllFields))

type FieldSet[F ~uint64] uint64

func (m FieldSet[F]) Includes(field F) bool {
	return F(m)&field != 0
}

// ListsBuilder for query
type ListsBuilder[F ~uint64] struct {
	Args  []any
	first int
	Names []string
	FieldSet[F]
	Fields F
}

func NewListsBuilder[F ~uint64](fieldsMask F) *ListsBuilder[F] {
	return &ListsBuilder[F]{
		FieldSet: FieldSet[F](fieldsMask),
	}
}

func (r *ListsBuilder[F]) AddField(field F, name string, val any) {
	if r.Includes(field) {
		r.Args = append(r.Args, val)
		if r.first == 0 {
			r.first = len(r.Args)
		}
		r.Names = append(r.Names, name)
		r.Fields |= field
	}
}

func (r *ListsBuilder[F]) NameList() string {
	return CommaList(r.Names)
}

func (r *ListsBuilder[F]) ValueListTemplate() string {
	return ValueListTemplate(r.first, len(r.Names))
}

func (r *ListsBuilder[F]) FieldListTemplate() string {
	return FieldListTemplate(r.first, r.Names)
}
