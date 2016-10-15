package functor

import (
	"github.com/arschles/functional-go/typeclass"
)

// IntegralFunctor is the typeclass for iterating over a container of Integral types, applying a function on each of them, and returning the results.
//
// All implementations of this interface must adhere to the following rules:
//
//	1. f.Map(func(i Integral) { return i }) == f
//	2. f.Map(funcA(funcB(param))) == f.Map(funcA).Map(funcB)
type IntegralFunctor interface {
	Map(func(typeclass.Integral) typeclass.Integral) IntegralFunctor
}
