package functor

import (
	"github.com/arschles/go-functional/util"
)

// OptionalIntFunctor is a container for an int that may or may not exist. You ma
type OptionalIntFunctor interface {
	util.Option
	// Map is the functor function. It applies fn to the contained int, if one exists
	Map(fn func(int) int) OptionalIntFunctor
	// Int returns the int that is contained. The return value of this function is undefined
	// if the functional.Option's Empty() method returns true
	Int() int
}

type optionalIntFunctorImpl struct {
	i      int
	exists bool
}

// NoInt returns an OptionalIntFunctor that is missing an int
func NoInt() OptionalIntFunctor {
	return optionalIntFunctorImpl{exists: false}
}

// SomeInt returns an OptionalIntFunctor that has an int
func SomeInt(i int) OptionalIntFunctor {
	return optionalIntFunctorImpl{exists: true, i: i}
}

func (o optionalIntFunctorImpl) Map(fn func(int) int) OptionalIntFunctor {
	if !o.exists {
		return o
	}
	o.i = fn(o.i)
	return o
}

func (o optionalIntFunctorImpl) Empty() bool {
	return !o.exists
}

func (o optionalIntFunctorImpl) Int() int {
	return o.i
}
