package functor

import (
	"fmt"

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

type optionalErrFunctorImpl struct {
	err error
}

// EmptyErr returns an OptionalIntFunctor that is missing an int
func EmptyErr() OptionalErrFunctor {
	return optionalErrFunctorImpl{err: nil}
}

// SomeErr returns an OptionalIntFunctor that has an int
func SomeErr(err error) OptionalErrFunctor {
	return optionalErrFunctorImpl{err: err}
}

func (o optionalErrFunctorImpl) Map(fn func(error) error) OptionalErrFunctor {
	if o.err != nil {
		o.err = fn(o.err)
		return o
	}
	return o
}

func (o optionalErrFunctorImpl) Empty() bool {
	return o.err == nil
}

func (o optionalErrFunctorImpl) Err() error {
	return o.err
}

func (o optionalErrFunctorImpl) String() string {
	if o.err == nil {
		return "empty"
	}
	return fmt.Sprintf("full(%s)", o.err.Error())
}
