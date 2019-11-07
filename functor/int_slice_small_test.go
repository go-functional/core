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

func TestIntSliceFunctorIntegrity(t *testing.T) {
	// chaining
	ints := intSlice(4)
	functor := LiftIntSlice(ints)
	m1 := functor.Map(plusOne).Map(timesTwo)

	// function composition
	ints = intSlice(4)
	functor = LiftIntSlice(ints)
	m2 := functor.Map(func(i int) int {
		return timesTwo(plusOne(i))
	})

	// compare the results
	assert.Equal(t, m1.Ints(), m2.Ints(), "returned slices are not equal")
}
