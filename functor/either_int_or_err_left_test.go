package functor

import (
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
}
