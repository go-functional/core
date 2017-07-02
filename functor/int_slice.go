package functor

// IntSliceFunctor is the typeclass for iterating over a slice of ints,
// applying a function on each of them, and returning the new IntSliceFunctor with the new results.
//
// All implementations of this interface must adhere to the following rules:
//
//	1. f.Map(func(i int) { return i }) == f
//		- this means that if you map with a function that does nothing (identity), you get the same
//			thing
//	2. f.Map(funcA(funcB(param))) == f.Map(funcA).Map(funcB)
//		- this means that you should be able to compose functions or execute them in serial
type IntSliceFunctor struct {
	ints []int
}

// ToIntSliceFunctor is the "constructor" to convert an []int into an IntSliceFunctor
func ToIntSliceFunctor(ints []int) IntSliceFunctor {
	return IntSliceFunctor{ints: ints}
}

// Map executes fn on every int in isf's internal slice and returns the resultant ints
func (isf IntSliceFunctor) Map(fn func(int) int) IntSliceFunctor {
	for i, elt := range isf.ints {
		retInt := fn(elt)
		isf.ints[i] = retInt
	}
	return isf
}

// Ints just returns a copy of the ints in isf
func (isf IntSliceFunctor) Ints() []int {
	return isf.ints
}
