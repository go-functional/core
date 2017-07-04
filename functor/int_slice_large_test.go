package functor

import (
	"testing"

	"github.com/arschles/assert"
)

func TestLargeIntSliceFunctor(t *testing.T) {
	ints := intSlice(10000)
	functor := LiftIntSlice(ints)

	mapped := functor.Map(plusOne).Map(minusOne)
	assert.Equal(t, len(mapped.Ints()), len(ints), "resultant ints not the same as original length")
	assert.Equal(t, mapped.Ints(), ints, "returned int slice")
}
