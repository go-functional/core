package channel

import (
	"context"

	"github.com/go-functional/core/values"
)

// import (
// 	"context"
// 	"fmt"

// 	"github.com/go-functional/core/values"
// )

// Join2 takes two channels and returns one that will either receive the values of
// the two, or receive a values.Error if one of the following happens:
//
// 1. ctx.Done() receives before both channels receive
// 2. Either channel is closed at any time
//
// THIS FUNCTION IS NOT YET IMPLEMENTED
func Join2[T, U any](ctx context.Context, ch1 <-chan T, ch2 <-chan U) <-chan values.Result[values.Tuple[T, U]] {
	type SuccessType = values.Tuple[T, U]
	retCh := make(chan values.Result[SuccessType])
	return retCh
	// TODO: implement
}
