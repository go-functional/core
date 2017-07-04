package functor

import (
	"github.com/go-functional/core/util"
)

// EitherIntOrErr is a container for either an int or an error. The left side of the
// either is the int, and the right side is an error. This interface is not a functor.
// Call ToLeft or ToRight to get functors that represent either side
type EitherIntOrErr interface {
	util.Either
	// ToLeft converts  is the functor function. It applies fn to the contained int, if one exists
	ToLeft() OptionalIntFunctor
	// Map is the functor function. It applies fn to the contained int, if
	ToRight() OptionalErrFunctor
}

type eitherIntOrErrImpl struct {
	i   int
	err error
}

// EitherIntOrErrLeft returns a left-side EitherIntOrErr (an int)
func EitherIntOrErrLeft(i int) EitherIntOrErr {
	return eitherIntOrErrImpl{i: i}
}

// EitherIntOrErrRight returns a right-side EitherIntOrErr (an error)
func EitherIntOrErrRight(err error) EitherIntOrErr {
	return eitherIntOrErrImpl{err: err}
}

// Left is the util.Either interface implementation. It returns true if e is an int
func (e eitherIntOrErrImpl) Left() bool {
	return e.err == nil
}

// Right is the util.Either interface implementation. It returns true if e is an error
func (e eitherIntOrErrImpl) Right() bool {
	return e.err != nil
}

func (e eitherIntOrErrImpl) ToLeft() OptionalIntFunctor {
	// TODO: implement
	return nil
}

func (e eitherIntOrErrImpl) ToRight() OptionalErrFunctor {
	// TODO: implement
	return nil
}
