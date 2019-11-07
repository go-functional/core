package functor

type parallelIntMapperResult struct {
	idx int
	val int
}

func parallelIntMapper(ints []int, fn func(int) int) []int {
	ch := make(chan parallelIntMapperResult)
	for i, elt := range ints {
		go func(idx int, elt int) {
			ch <- parallelIntMapperResult{idx: idx, val: fn(elt)}
		}(i, elt)
	}
	for range ints {
		elt := <-ch
		ints[elt.idx] = elt.val
	}
	return ints
}
