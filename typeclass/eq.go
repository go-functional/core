package typeclass

import (
	"strconv"
)

// Eq is the equality typeclass: (Eq a) => a -> a -> Bool
type Eq interface {
	Eq(Eq) bool
}

// IntEq is a basic type class to check if an int and other Eq types are equal
type IntEq int

// Eq is the Eq interface implementation. It checks for equality with other types
func (i IntEq) Eq(e Eq) bool {
	it := int(i)
	switch x := e.(type) {
	case IntEq:
		return it == int(x)
	case StrEq:
		otherI, err := strconv.Atoi(string(x))
		if err != nil {
			return false
		}
		return it == otherI
	default:
		return false
	}
}

// StrEq is the Eq type class implementation for strings
type StrEq string

// Eq is the Eq interface implementation
func (s StrEq) Eq(e Eq) bool {
	st := string(s)
	switch x := e.(type) {
	case StrEq:
		return st == string(x)
	case IntEq:
		return st == strconv.Itoa(int(x))
	default:
		return false
	}
}

// IntSliceEq is the Eq type class implementation for int slices
type IntSliceEq []int

// Eq determines if s is equal to e
func (s IntSliceEq) Eq(e Eq) bool {
	switch x := e.(type) {
	case IntSliceEq:
		return intSlicesEq(s, x)
	default:
		return false
	}
}

func intSlicesEq(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, elt1 := range sl1 {
		elt2 := sl2[i]
		if elt1 != elt2 {
			return false
		}
	}
	return true
}
