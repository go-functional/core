package functor

import (
	"testing"
)

func TestIntSliceFunctor(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	functor := IntSliceFunctor{ints: ints}
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
	retFunctor := functor.Map(plusOne).Map(log).Map(minusOne)
	retIntSliceFunctor, ok := retFunctor.(IntSliceFunctor)
	if !ok {
		t.Fatalf("returned functor was not an IntegralFunctor")
	}
	if len(retIntSliceFunctor.ints) != len(ints) {
		t.Fatalf(
			"resultant ints length (%d) is not the same as original length (%d)",
			len(retIntSliceFunctor.ints),
			len(ints),
		)
	}
	for i, origInt := range ints {
		retInt := retIntSliceFunctor.ints[i]
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
