package functor

import (
	"errors"
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
}
