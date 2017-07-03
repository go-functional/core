package functor

import (
	"testing"

	"github.com/go-functional/core/typeclass"
)

func TestIntegralSliceFunctor(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	functor := IntegralSliceFunctor{ints: ints}
	plusOne := func(i typeclass.Integral) typeclass.Integral {
		return typeclass.IntIntegral(i.Int() + 1)
	}
	minusOne := func(i typeclass.Integral) typeclass.Integral {
		return typeclass.IntIntegral(i.Int() - 1)
	}
	log := func(i typeclass.Integral) typeclass.Integral {
		t.Logf("%#v", i)
		return i
	}
	retFunctor := functor.Map(plusOne).Map(log).Map(minusOne)
	retIntSliceFunctor, ok := retFunctor.(IntegralSliceFunctor)
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
