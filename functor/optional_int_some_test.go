package functor

import (
	"testing"
)

func TestSomeInt(t *testing.T) {
	const i = 1
	mapFn := func(i int) int {
		return i + 4
	}
	si := SomeInt(i)
	if si.Empty() {
		t.Fatalf("SomeInt reported as empty")
	}
	if si.Int() != 1 {
		t.Fatalf("SomeInt.Int() reported %d, expected %d", si.Int(), i)
	}
	newSI := si.Map(mapFn)
	if newSI.Empty() {
		t.Fatalf("SomeInt reported as empty")
	}
	if newSI.Int() != mapFn(i) {
		t.Fatalf("SomeInt.Int() reported %d, expected %d", newSI.Int(), mapFn(i))
	}
}
