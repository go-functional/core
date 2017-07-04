package functor

import (
	"errors"
	"testing"

	"github.com/arschles/assert"
)

func TestEmptyErr(t *testing.T) {
	// this is the function that we'll pass to Map
	mapFunc := func(error) error {
		return errors.New("other error")
	}

	// create an error that doesn't exist
	emptyErr := EmptyErr()
	assert.NoErr(t, emptyErr.Err())
	assert.True(t, emptyErr.Empty(), "expected EmptyErr to return empty")

	mapped := emptyErr.Map(mapFunc)
	assert.NoErr(t, mapped.Err())
	assert.True(t, mapped.Empty(), "expected mapped EmptyErr to return empty")
}
