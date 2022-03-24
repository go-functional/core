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

// ToSlice returns n as a slice. Note that every invocation
// of this function will create a _new_ slice
func (n *Nonempty[T]) ToSlice() []T {
	ret := append([]T{n.head}, n.rest...)
	return ret
}

// Head returns the head of n.
// Due to the invariant that the list is non-empty,
// this function always returns a valid value.
func (n *Nonempty[T]) Head() T {
	return n.head
}

// Tail returns all the elements of n that come after
// the head. If n has only 1 total elements, this returns
// nil
func (n *Nonempty[T]) Tail() []T {
	return n.rest
}

// Append appends vals to the end of n
func (n *Nonempty[T]) Append(vals ...T) {
	n.rest = append(n.rest, vals...)
}
