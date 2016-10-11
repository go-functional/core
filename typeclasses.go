package funcgo

// see http://learnyouahaskell.com/types-and-typeclasses

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
