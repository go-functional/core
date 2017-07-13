package functor

import (
	"fmt"
)

// LiftIntSliceFromChan converts a readonly chan of ints into an
// IntSliceFunctor. In FP terms, this is called "lifting".
//
// The functor's Map function will block until the channel is closed, and the
// function that you pass to map will be called on each int that is received
// in order.
//
// Note that the Int() function will not return any values until Map has been
// called
func LiftIntSliceFromChan(ch <-chan int) IntSliceFunctor {
	return intSliceFromChanFunctorImpl{ch: ch}
}

type intSliceFromChanFunctorImpl struct {
	ch    <-chan int
	slice []int
}

func (i intSliceFromChanFunctorImpl) Map(fn func(int) int) IntSliceFunctor {
	slc := []int{}
	for elt := range i.ch {
		slc = append(slc, fn(elt))
	}
	i.slice = slc
	return i
}

func (i intSliceFromChanFunctorImpl) Ints() []int {
	return i.slice
}

func (i intSliceFromChanFunctorImpl) String() string {
	return fmt.Sprintf("%#v", i.slice)
}
