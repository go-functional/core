package functor

import (
	"errors"
	"testing"
)

func TestEmptyErr(t *testing.T) {
	// this is the function that we'll pass to Map
	mapFunc := func(error) error {
		return errors.New("other error")
	}

	// create an error that doesn't exist
	emptyErr := EmptyErr()
	if emptyErr.Err() != nil {
		t.Fatalf("expected no error on EmptyErr")
	}
	if !emptyErr.Empty() {
		t.Fatalf("expected EmptyErr to return empty")
	}

	mapped := emptyErr.Map(mapFunc)
	if mapped.Err() != nil {
		t.Fatalf("expected mapped functor's error to be nil")
	}
	if !mapped.Empty() {
		t.Fatalf("expected mapped EmptyErr to not be empty")
	}
}
