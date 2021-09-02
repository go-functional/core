package iter


func FlatMap[T any](slc []T, fn func(t T) []U) []U {
	ret := []U{}
	for i, val := range slc {
		ret = append(ret, fn(val)...)
	}
	return ret
}

