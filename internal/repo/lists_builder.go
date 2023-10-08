package repo

import (
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

// ListsBuilder for insert/update query
type ListsBuilder struct {
	Args  []any
	first int
	Names []string
}

func (r *ListsBuilder) AddField(enable bool, name string, val any) {
	if enable {
		r.Args = append(r.Args, val)
		if r.first == 0 {
			r.first = len(r.Args)
		}
		r.Names = append(r.Names, name)
	}
}

func (r *ListsBuilder) NameList() string {
	return CommaList(r.Names)
}

func (r *ListsBuilder) ValueListTemplate() string {
	return ValueListTemplate(r.first, len(r.Names))
}

func (r *ListsBuilder) FieldListTemplate() string {
	return FieldListTemplate(r.first, r.Names)
}
