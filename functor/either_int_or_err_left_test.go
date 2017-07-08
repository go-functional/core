package functor

import (
	"errors"
	"testing"

	"github.com/arschles/assert"
)

func TestEitherIntOrErrLeft(t *testing.T) {
	const i = 123
	left := EitherIntOrErrLeft(i)
	assert.True(t, left.Left(), "Left was not reported as a left")
	assert.False(t, left.Right(), "Left was reported as a right")
	leftProj := left.ToLeft()
	rightProj := left.ToRight()
	assert.False(t, leftProj.Empty(), "left projection reported as empty")
	assert.True(t, rightProj.Empty(), "right projection not reported as empty")
	assert.Equal(t, leftProj.Int(), i, "left projection int")

	// mapping over the left projection should do something
	intFunc := func(i int) int {
		return i + 1
	}
	mappedLeft := leftProj.Map(intFunc)
	assert.False(t, mappedLeft.Empty(), "left mapped projection reported as empty")
	assert.Equal(t, mappedLeft.Int(), intFunc(i), "returned int")

	// mapping over the right projection should do nothing
	mappedRight := rightProj.Map(func(error) error {
		return errors.New("neverShouldBeReturned")
	})
	assert.True(t, mappedRight.Empty(), "right mapped projection reported as not empty")
	// don't check the error return - it is undefined if Empty is returned as true
}
