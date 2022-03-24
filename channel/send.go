package channel

import "context"

func SendWithContext[T any](
	ctx context.Context,
	ch chan<- T,
	val T,
) bool {
	select {
	case <-ctx.Done():
		return false
	case ch <- val:
		return true
	}
}
