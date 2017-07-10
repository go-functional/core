package functor

import (
	"fmt"

	"github.com/go-functional/core/util"
)

// EitherIntOrErr is a container for either an int or an error. The left side of the
// either is the int, and the right side is an error. This interface is not a functor.
// Call ToLeft or ToRight to get functors that represent either side
type EitherIntOrErr interface {
	util.Either
	// ToLeft converts the either to the left side. The resulting OptionalIntFunctor
	// will be full if this either was the left side, and empty if right
	ToLeft() OptionalIntFunctor
	// ToRight converts the either to the right side. The resulting OptionalErrFunctor
	// will be full if this either was the right side, and empty if left
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

// ToLeft returns an optional int that will be Some if e.Left returns true.
// In FP terms, this is a called a "Left projection"
func (e eitherIntOrErrImpl) ToLeft() OptionalIntFunctor {
	if e.err == nil {
		return SomeInt(e.i)
	}
	return EmptyInt()
}

// ToRight returns an optional err that will be Some if e.Right returns true.
// In FP terms, this is called a "Right projection"
func (e eitherIntOrErrImpl) ToRight() OptionalErrFunctor {
	if e.err != nil {
		return SomeErr(e.err)
	}
	return EmptyErr()
}

func (e eitherIntOrErrImpl) String() string {
	if e.Left() {
		return fmt.Sprintf("left(%d)", e.i)
	}
	return fmt.Sprintf("right(%s)", e.err)
}
