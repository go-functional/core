package list

// Nonempty is a list that is guaranteed to always
// have at least 1 element
type Nonempty[T any] struct {
	head T
	rest []T
}

// Newnonempty creates a new Nonempty list with the
// given head and optional tail
func NewNonempty[T any](head T, tail ...T) Nonempty[T] {
	return Nonempty[T]{
		head: head,
		rest: tail,
	}
}

func (n *Nonempty[T]) toSlice() []T {
	ret := append([]T{n.head}, n.rest...)
	return ret
}

func (n *Nonempty[T]) Head() T {
	return n.head
}

func (n *Nonempty[T]) Tail() []T {
	return n.rest
}
