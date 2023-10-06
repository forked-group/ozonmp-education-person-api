package repo

import (
	"strconv"
	"strings"
)

func anySlice[T any](a []T) []any {
	res := make([]any, len(a))

	for i := range res {
		res[i] = a[i]
	}

	return res
}

func placeholderList(first int, n int) string {
	elems := make([]string, n)

	for i, j := 0, first; i < n; i, j = i+1, j+1 {
		elems[i] = "$" + strconv.Itoa(j)
	}

	return strings.Join(elems, ",")
}

func placeholderSetList(first int, names []string) string {
	n := len(names)
	elems := make([]string, n)

	for i, j := 0, first; i < n; i, j = i+1, j+1 {
		elems[i] = names[i] + "=$" + strconv.Itoa(j)
	}

	return strings.Join(elems, ",")
}

type field interface {
	~uint64 | ~int64 | ~uint | ~int
}

type params[F field] struct {
	names  []string
	values []any
	mask   F
	fields F
}

func newParamsWithMask[F field](mask F) *params[F] {
	return &params[F]{mask: mask}
}

func (p *params[F]) add(field F, name string, value any) {
	if p.mask&field != 0 {
		p.names = append(p.names, name)
		p.values = append(p.values, value)
		p.fields |= field
	}
}
