package functor

import (
	"sync"
)

type parallelIntMapperResult struct {
	idx int
	val int
}

func parallelIntMapper(ints []int, fn func(int) int) []int {
	resultsCh := make(chan parallelIntMapperResult)
	var wg sync.WaitGroup
	for i, elt := range ints {
		wg.Add(1)
		ch := make(chan parallelIntMapperResult)
		go func(idx int, elt int) {
			ch <- parallelIntMapperResult{idx: idx, val: fn(elt)}
		}(i, elt)
		go func() {
			defer wg.Done()
			elt := <-ch
			resultsCh <- elt
		}()
	}
	go func() {
		wg.Wait()
		close(resultsCh)
	}()
	for elt := range resultsCh {
		ints[elt.idx] = elt.val
	}
	return ints
}
