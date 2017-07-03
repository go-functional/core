package functor

import (
	"github.com/go-functional/core/util"
)

// OptionalErrFunctor is a container for an error that may or may not exist.
type OptionalErrFunctor interface {
	util.Option
	// Map is the functor function. It applies fn to the contained int, if one exists
	Map(fn func(error) error) OptionalErrFunctor
	// Err returns the int that is contained. The return value of this function is undefined
	// if the Option's Empty() method returns true
	Err() error
}
