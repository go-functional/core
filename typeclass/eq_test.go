package typeclass

import (
	"testing"

	"github.com/arschles/assert"
)

// This is how we effectively can do "1" == 1
func TestIntEqStrEq(t *testing.T) {
	intEq := IntEq(1)
	strEq := StrEq("1")
	assert.True(t, intEq.Eq(strEq), "1 != '1'")
	assert.True(t, strEq.Eq(intEq), "'1' != 1")
}
