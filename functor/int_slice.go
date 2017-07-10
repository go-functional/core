package functor

import (
	"fmt"
)

// IntSliceFunctor is a container of []int, and a facility for easily iterating over a slice of ints,
// applying a function on each of them, and returning the new IntSliceFunctor with the new results.
//
// All implementations of this interface must adhere to the following rules:
//
//	1. f.Map(func(i int) { return i }) == f
//		- this means that if you map with a function that does nothing (identity), you get the same
//			thing
//	2. f.Map(funcA(funcB(param))) == f.Map(funcA).Map(funcB)
//		- this means that you should be able to compose functions or execute them in serial
type IntSliceFunctor interface {
	fmt.Stringer
	// Map is the Functor function. It applies fn to every element in the contained slice
	Map(fn func(int) int) IntSliceFunctor
	// Ints is just a convenience function to get the int slice that the functor holds
	Ints() []int
}

type intSliceFunctorImpl struct {
	ints []int
}

// LiftIntSlice converts an int slice into an IntSliceFunctor. In FP terms, this operation
// is called "lifting", and in many languages it's done automatically
func LiftIntSlice(slice []int) IntSliceFunctor {
	return intSliceFunctorImpl{ints: slice}
}

// Map executes fn on every int in isf's internal slice and returns the resultant ints
func (isf intSliceFunctorImpl) Map(fn func(int) int) IntSliceFunctor {
	if len(isf.ints) < 100 {
		isf.ints = serialIntMapper(isf.ints, fn)
		return isf
	}
	isf.ints = parallelIntMapper(isf.ints, fn)
	return isf
}

// Ints just returns a copy of the ints in isf
func (isf intSliceFunctorImpl) Ints() []int {
	return isf.ints
}

// String just returns the string representation of the functor.
//
// This is the implementation of fmt.Stringer
func (isf intSliceFunctorImpl) String() string {
	return fmt.Sprintf("%#v", isf.ints)
}
