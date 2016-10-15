package typeclass

import (
	"strconv"
)

// Show is the typeclass for converting a type into a string: (Show a) => a -> string
// This is similar to the fmt.Stringer interface, but it obviously is implemented outside of the type. A common use case would be to have many of these, for different purposes
type Show interface {
	Show() string
}

// IntShow is the Show instance for ints
type IntShow int

// Show is the Show interface implementation
func (i IntShow) Show() string {
	return strconv.Itoa(int(i))
}

// StringShow is the show instance for strings. It's trivial, and just here for composition purposes
type StringShow string

// Show is the Show interface implementation. This is the identity function
func (s StringShow) Show() string {
	return string(s)
}
