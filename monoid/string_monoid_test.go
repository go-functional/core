package monoid

import (
	"testing"

	"github.com/arschles/assert"
)

func TestStringMonoid(t *testing.T) {
	const str = "abcde"
	monoid := LiftStringMonoid(str)
	assert.Equal(t, monoid.Zero(), "", "zero val")
	assert.Equal(t, monoid.String(), str, "escape hatch value")
	appended := monoid.Append(str)
	assert.Equal(t, appended.String(), str+str, "appended escape hatch val")
}
