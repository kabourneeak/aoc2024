package days

func Filter[T any](source []T, predicate func(T) bool) (out []T) {
	for _, s := range source {
		if predicate(s) {
			out = append(out, s)
		}
	}
	return
}

func Swap[T any](source []T, i, j int) {
	source[j], source[i] = source[i], source[j]
}
