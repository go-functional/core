package functor

import (
	"github.com/go-functional/core/typeclass"
)

// IntegralFunctor is the typeclass for iterating over a container of Integral typeclasses,
// applying a function on each of them, and returning the results.
//
// All implementations of this interface must adhere to the following rules:
//
//	1. f.Map(func(i Integral) { return i }) == f
//	2. f.Map(funcA(funcB(param))) == f.Map(funcA).Map(funcB)
type IntegralFunctor interface {
	Map(func(typeclass.Integral) typeclass.Integral) IntegralFunctor
}

// IntegralSliceFunctor is the IntegralFunctor implementation for []int types
type IntegralSliceFunctor struct {
	ints []int
}

// Map is the IntegralFunctor interface implementation
func (isf IntegralSliceFunctor) Map(fn func(typeclass.Integral) typeclass.Integral) IntegralFunctor {
	for i, elt := range isf.ints {
		retIntegral := fn(typeclass.IntIntegral(elt))
		isf.ints[i] = int(retIntegral.Int())
	}
	return isf
}
