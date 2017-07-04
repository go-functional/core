package functor

import (
	"testing"

	"github.com/arschles/assert"
)

func TestEmptyInt(t *testing.T) {
	mapFn := func(i int) int {
		return 0
	}
	ei := EmptyInt()
	assert.True(t, ei.Empty(), "EmptyInt not reported as empty")

	mapped := ei.Map(mapFn)
	assert.True(t, mapped.Empty(), "mapped EmptyInt not reported as empty")
}
