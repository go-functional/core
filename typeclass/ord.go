package typeclass

import (
	"strconv"
)

// Ord is the Ordering typeclass: (Ord a) => a -> a -> Bool
//
// This type is vaguely similar to the Interface interface in the sort package
type Ord interface {
	// Before returns true if the current is strictly less than the parameter
	Before(Ord) bool
}

// IntOrd is the Ord implementation for ints
type IntOrd int

// Before is the Ord interface implementation
func (i IntOrd) Before(o Ord) bool {
	it := int(i)
	switch x := o.(type) {
	case IntOrd:
		return it < int(x)
	case StringOrd:
		otherI, err := strconv.Atoi(string(x))
		if err != nil {
			return false
		}
		return it < otherI
	default:
		return false
	}
}

// StringOrd is the Ord implementation for strings
type StringOrd string

// Before is the interface implementation for strings
func (s StringOrd) Before(o Ord) bool {
	st := string(s)
	switch x := o.(type) {
	case IntOrd:
		return st < strconv.Itoa(int(x))
	case StringOrd:
		return st < string(x)
	default:
		return false
	}
}
