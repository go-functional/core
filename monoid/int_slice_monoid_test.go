package monoid

import (
	"testing"

	"github.com/arschles/assert"
	"github.com/go-functional/core/typeclass"
)

func TestIntSlice(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	m := LiftIntSlice(ints)
	zeroEq := typeclass.IntSliceEq(m.Zero())
	expectedEq := typeclass.IntSliceEq(nil)
	assert.True(t, zeroEq.Eq(expectedEq), "zero value was not empty")
	assert.Equal(t, m.Ints(), ints, "escape hatch")
	appended := m.Append(ints...)
	assert.Equal(t, appended.Ints(), append(ints, ints...), "appended")
}
