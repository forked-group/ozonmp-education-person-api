package repo

func anySlice[T any](a []T) []any {
	res := make([]any, len(a))

	for i := range res {
		res[i] = a[i]
	}

	return res
}
