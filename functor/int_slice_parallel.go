package functor

import (
	"sync"
)

type parallelIntMapperResult struct {
	idx int
	val int
}

func parallelIntMapper(ints []int, fn func(int) int) []int {
	ch := make(chan parallelIntMapperResult)
	var wg sync.WaitGroup
	for i, elt := range ints {
		wg.Add(1)
		go func(idx int, elt int) {
			defer wg.Done()
			ch <- parallelIntMapperResult{idx: idx, val: fn(elt)}
		}(i, elt)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for elt := range ch {
		ints[elt.idx] = elt.val
	}
	return ints
}
