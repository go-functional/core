package channel

import "context"

// Filter receives every value on ch, passes it to fn, and
// if fn returns true, forwards that value to the returned
// channel. The returned channel is closed if ctx.Done()
// receives or ch is closed.
func Filter[T any](
	ctx context.Context,
	ch <-chan T,
	fn func(T) bool,
) <-chan T {
	ret := make(chan T)
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
				shouldSend := fn(val)
				if !shouldSend {
					continue
				}
				if SendWithContext(ctx, ret, val) != nil {
					return
				}
			}
		}
	}()
	return ret
}
