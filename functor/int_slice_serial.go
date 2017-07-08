package functor

func serialIntMapper(ints []int, fn func(int) int) []int {
	for i, elt := range ints {
		retInt := fn(elt)
		ints[i] = retInt
	}
	return ints
}
