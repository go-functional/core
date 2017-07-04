package functor

import (
	"testing"
)

func TestEmptyInt(t *testing.T) {
	mapFn := func(i int) int {
		return 0
	}
	ei := EmptyInt()
	if !ei.Empty() {
		t.Fatalf("EmptyInt not reported as empty")
	}
	mapped := ei.Map(mapFn)
	if !mapped.Empty() {
		t.Fatalf("mapped EmptyInt not reported as empty")
	}
}
