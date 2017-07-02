package functor

import (
	"testing"
)

func TestIntSliceFunctor(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	functor := LiftIntSlice(ints)
	plusOne := func(i int) int {
		return i + 1
	}
	minusOne := func(i int) int {
		return i - 1
	}
	log := func(i int) int {
		t.Logf("%#v", i)
		return i
	}
	retIntSliceFunctor := functor.Map(plusOne).Map(log).Map(minusOne)
	if len(retIntSliceFunctor.Ints()) != len(ints) {
		t.Fatalf(
			"resultant ints length (%d) is not the same as original length (%d)",
			len(retIntSliceFunctor.Ints()),
			len(ints),
		)
	}
	for i, origInt := range ints {
		retInt := retIntSliceFunctor.Ints()[i]
		if origInt != retInt {
			t.Fatalf(
				"expected %d for int # %d, got %d instead",
				origInt,
				i,
				retInt,
			)
		}
	}
}
