package channel

import (
	"context"
	"fmt"

	"github.com/go-functional/core/values"
)

func RecieveWithContext[T any](ctx context.Context, ch <-chan T) (T, error) {
	var zero T
	select {
	case <-ctx.Done():
		return zero, ctx.Err()
	case val, ok := <-ch:
		if !ok {
			return zero, fmt.Errorf("channel closed")
		}
		return val, nil
	}
}

// ReceiveChanWithContext listens on ch in the background and returns a channel
// that follows the following rules:
//
// 1. If ctx.Done() receives before ch does, the returned channel will receive a
// values.Error and then be closed
// 2. If ch is closed before anything else, the returned channel will receive
// a values.Error and then be closed
// 3. If ch receives a value, the returned channel will receive that value and then
// be closed
func ReceiveChanWithContext[T any](ctx context.Context, ch <-chan T) <-chan values.Result[T] {
	retCh := make(chan values.Result[T])
	go func() {
		defer close(retCh)
		select {
		case <-ctx.Done():
			retCh <- *values.Error[T](ctx.Err())
		case val, ok := <-ch:
			if !ok {
				retCh <- *values.Error[T](fmt.Errorf("channel closed"))
			} else {
				retCh <- *values.Ok(val)
			}
		}
	}()
	return retCh
}
