package functor

import (
	"errors"
	"testing"

	"github.com/arschles/assert"
)

func TestSomeErr(t *testing.T) {
	testErr := errors.New("tester")
	otherTestErr := errors.New("othertesterr")

	// this is the function that we'll pass to Map
	mapFunc := func(error) error {
		return otherTestErr
	}

	// create an error that exists
	someErr := SomeErr(testErr)
	assert.Err(t, testErr, someErr.Err())
	assert.False(t, someErr.Empty(), "expected SomeErr to return not empty")

	mapped := someErr.Map(mapFunc)
	assert.Err(t, mapFunc(testErr), mapped.Err())
	assert.False(t, mapped.Empty(), "expected mapped SomeErr to not be empty")
}
