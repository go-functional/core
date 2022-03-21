package slice

func FlatMap[T, U any](slc []T, fn func(t T) []U) []U {
	ret := []U{}
	for _, val := range slc {
		ret = append(ret, fn(val)...)
	}
	return ret
}
