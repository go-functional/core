package funcgo

// see http://learnyouahaskell.com/types-and-typeclasses

// Each of the typeclasses herein are "adapters" for existing types. Other languages (Haskell, Scala) automatically convert existing types for use with these adapters, but in Go we have to do things more manually. Here's how:
//
// 1. create an implementation of one of the type classes
//
//	type intEq int
//
//	func (i intEq) Eq(e Eq) bool {
//		i2, ok := e.(int)
//		if !ok {
//			return false
//		}
//		return i == i2
//	}
//
// 2. wrap your value in the implementation
//
//	intEq(1).Eq(intEq(2))

// Eq is the equality typeclass: (Eq a) => a -> a -> Bool
type Eq interface {
	Eq(Eq) bool
}

// Ord is the Ordering typeclass: (Ord a) => a -> a -> Bool
type Ord interface {
	Before(Ord) bool
}

// Show is the typeclass for converting a type to a string (same as fmt.Stringer): (Show a) => a -> string
type Show interface {
	Show() string
}

// Read is the typeclass for converting a string into an instance of Read: (Read a) => string -> Read
type Read interface {
	Read(string) (Read, error)
}

// Integral is the typeclass for representing an integer type
type Integral interface {
	Int() int64
}

// IntegralFunctor is the typeclass for iterating over a container of Integral types, applying a function on each of them, and returning the results.
//
// All implementations of this interface must adhere to the following rules:
//
//	1. f.Map(func(i Integral) { return i }) == f
//	2. f.Map(funcA(funcB(param))) == f.Map(funcA).Map(funcB)
type IntegralFunctor interface {
	Map(func(Integral) Integral) IntegralFunctor
}

// Floating is the typeclass for representing a floating type
type Floating interface {
	Float() float64
}
