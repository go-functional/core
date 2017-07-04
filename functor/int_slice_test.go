package functor

import (
	"testing"

	"github.com/arschles/assert"
)

func intSlice(numInts int) []int {
	slice := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		slice[i] = i + 5
	}
	return slice
}
func TestIntSliceFunctor(t *testing.T) {
	testCases := [][]int{
		intSlice(4),
		intSlice(10000),
	}
	for _, testCase := range testCases {
		ints := testCase
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
		assert.Equal(
			t,
			len(retIntSliceFunctor.Ints()),
			len(ints),
			"resultant ints not the same as original length",
		)

		for i, origInt := range ints {
			retInt := retIntSliceFunctor.Ints()[i]
			assert.Equal(t, retInt, origInt, "returned int")
		}
	}
}
