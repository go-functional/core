// Package typeclass contains example type classes
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
//
// See http://learnyouahaskell.com/types-and-typeclasses for more theoretical details on how and why these work
package typeclass
