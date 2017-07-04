package functor

import (
	"fmt"
	"testing"

	"github.com/arschles/assert"
)

func TestSomeInt(t *testing.T) {
	const i = 1
	mapFn := func(i int) int {
		return i + 4
	}
	si := SomeInt(i)
	assert.False(t, si.Empty(), "SomeInt reported as empty")
	assert.Equal(t, si.Int(), i, fmt.Sprintf(
		"SomeInt.Int() reported %d, expected %d",
		si.Int(),
		i,
	))

	mapped := si.Map(mapFn)
	assert.False(t, mapped.Empty(), "mapped SomeInt reported as empty")
	assert.Equal(
		t,
		mapped.Int(),
		mapFn(i),
		fmt.Sprintf(
			"mapped SomeInt.Int() reported %d, expected %d",
			mapped.Int(),
			mapFn(i),
		),
	)
}
