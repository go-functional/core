package functor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/arschles/assert"
)

func TestEitherIntOrErrRight(t *testing.T) {
	err := errors.New("testerr")
	right := EitherIntOrErrRight(err)
	assert.False(t, right.Left(), "Right was reported as a left")
	assert.True(t, right.Right(), "right was not reported as a right")
	leftProj := right.ToLeft()
	rightProj := right.ToRight()
	assert.True(t, leftProj.Empty(), "left projection not reported as empty")
	assert.False(t, rightProj.Empty(), "right projection reported as empty")
	assert.Equal(t, rightProj.Err(), err, "right projection error")

	errFunc := func(err error) error {
		return fmt.Errorf("%s1", err.Error())
	}
	mappedRight := rightProj.Map(errFunc)
	assert.False(t, mappedRight.Empty(), "right mapped projection reported as empty")
	assert.Equal(t, mappedRight.Err(), errFunc(err), "returned error")

	mappedLeft := leftProj.Map(func(int) int {
		// this function should never be called
		return 1
	})
	assert.True(t, mappedLeft.Empty(), "left mapped projection reported as not empty")
	// don't check the Int method - its return value is undefined if Empty returns true
}
