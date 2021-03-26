package iter

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// Iter iterates from 0 through n-1 (not including n) calling fn on each iteration,
// and passing the iteration number each time.
//
// If fn returns an error, Iter returns with that error immediately.
//
// Common usage of this function is to iterate through a slice. See example below:
//
//	slc := []int{1, 2, 3, 4, 5}
//	Iter(len(slc), func(i uint) error {
//		slc[i] += 1
//	})
func Iter(n uint, fn func(i uint) error) error {
	for i := uint(0); i < n; i++ {
		if err := fn(i); err != nil {
			return err
		}
	}
	return nil
}

func errListFunc(n uint, fn func(uint) error) func() error {
	return func() error {
		return fn(n)
	}
}

// IterPar calls fn the given n number of times, each timne in parallel.
//
// The function will be called in parallel n times, but if any one of them
// returns an error, that error will be returned from this function.
// If more than one of them returns an error, the first one to finish
// will be returned from this function. That behavior means that the
// return value of this function may be non-deterministic.
//
// Common use of this function is to do expensive operations on a slice
// in parallel. Note that the below example reads and writes shared
// state concurrently, so it protects those accesses with a mutex
//
//	var mut sync.Mutex
//	slc := []int{1, 2, 3, 4, 5}
//	IterPar(context.Background(), len(slc), func(i uint) error {
//		mut.Lock()
//		defer mut.Unlock()
//		newVal, err := expensiveOperation(slc[i])
//		if err != nil {
//			return err
//		}
//		slc[i] = newVal
//		return nil
//	})
func IterPar(ctx context.Context, n uint, fn func(i uint) error) error {
	errgrp, _ := errgroup.WithContext(ctx)
	for i := uint(0); i < n; i++ {
		errgrp.Go(errListFunc(i, fn))
	}
	return errgrp.Wait()
}
