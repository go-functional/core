package functor

import (
	"testing"

	"github.com/arschles/assert"
)

func TestSmallIntSliceFunctor(t *testing.T) {
	ints := intSlice(4)
	functor := LiftIntSlice(ints)

	mapped := functor.Map(plusOne).Map(minusOne)
	assert.Equal(t, len(mapped.Ints()), len(ints), "resultant ints not the same as original length")
	assert.Equal(t, mapped.Ints(), ints, "returned int slice")
}
