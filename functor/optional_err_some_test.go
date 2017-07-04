package functor

import (
	"errors"
	"testing"
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
	if someErr.Err() != testErr {
		t.Fatalf("expected error %s, got %s", testErr, someErr.Err())
	}
	if someErr.Empty() {
		t.Fatalf("expected SomeErr to return not empty")
	}

	mapped := someErr.Map(mapFunc)
	if mapped.Err() != mapFunc(testErr) {
		t.Fatalf(
			"expected mapped functor's error to be %s, got %s",
			mapFunc(testErr),
			mapped.Err(),
		)
	}
	if mapped.Empty() {
		t.Fatalf("expected mapped SomeErr to not be empty")
	}
}

