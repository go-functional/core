package list

func Reduce[T, U any](
	slc []T,
	start U,
	fn func(U, T) U,
) U {
	ret := start
	for _, t := range slc {
		ret = fn(ret, t)
	}
	return ret
}
