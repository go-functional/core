package monoid

import (
	"fmt"
)

// StringMonoid is the monoid container for strings
type StringMonoid interface {
	fmt.Stringer
	Zero() string
	Append(s string) StringMonoid
}

// LiftStringMonoid lifts a string into the monoid container
func LiftStringMonoid(s string) StringMonoid {
	return stringMonoidImpl{str: s}
}

type stringMonoidImpl struct {
	str string
}

func (s stringMonoidImpl) Append(str string) StringMonoid {
	s.str = s.str + str
	return s
}

func (stringMonoidImpl) Zero() string {
	return ""
}

func (s stringMonoidImpl) String() string {
	return s.str
}
