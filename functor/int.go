package functor

// IntFunctor is the typeclass for iterating over a container of ints,
// applying a function on each of them, and returning the new IntFunctor with the new results.
//
// All implementations of this interface must adhere to the following rules:
//
//	1. f.Map(func(i int) { return i }) == f
//		- this means that if you map with a function that does nothing (identity), you get the same
//			thing
//	2. f.Map(funcA(funcB(param))) == f.Map(funcA).Map(funcB)
//		- this means that you should be able to compose functions or execute them in serial
type IntFunctor interface {
	Map(func(int) int) IntFunctor
}

// IntSliceFunctor is the IntegralFunctor implementation for []int types
type IntSliceFunctor struct {
	ints []int
}

// Map is the IntFunctor interface implementation
func (isf IntSliceFunctor) Map(fn func(int) int) IntFunctor {
	for i, elt := range isf.ints {
		retInt := fn(elt)
		isf.ints[i] = retInt
	}
	return isf
}
