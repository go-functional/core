package channel

import "context"

// Map starts listens on ch in a background goroutine.
// Upon receiving a value, it calls fn(receivedValue)
// and sends the result to the channel this function
// returns.
//
// If ch is closed or ctx.Done() receives, this function
// stops receiving on ch and closes the returned channel.
func Map[T, U any](
	ctx context.Context,
	ch <-chan T,
	fn func(T) U,
) <-chan U {
	ret := make(chan U)
	select {
	case <-ctx.Done():
		close(ret)
		return ret
	default:
	}
	go func() {
		defer close(ret)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-ch:
				if !ok {
					return
				}
				if !SendWithContext(ctx, ret, fn(val)) {
					return
				}

			}
		}
	}()
	return ret
}
