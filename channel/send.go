package channel

import "context"

// SendWithContext either sends val to ch and returns nil
// or, if ctx.Done() receives, returns ctx.Err()
func SendWithContext[T any](
	ctx context.Context,
	ch chan<- T,
	val T,
) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case ch <- val:
		return nil
	}
}
