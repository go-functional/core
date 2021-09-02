package slice

import "errors"

// Cons creates a new list with head at the front of tail
func Cons[T any](head T, tail []T) []T {
	lst := []T{head}
	return append(lst, tail...)
}

// Head returns slc[0] if the list has at least one element
// in it. Otherwise, returns empty() and a descriptive, non-nil
// error
func Head[T any](slc []T, empty func() T) (T, error) {
	if len(slc) == 0 {
		return empty(), errors.New("Head called on empty list")
	}
	return slc[0], nil
}

// Tail returns slc[1:] only if the list has 2 or more elements in
// it. If it doesn't, returns an empty list and a descriptive, non-nil
// error
func Tail[T any](slc []T) ([]T, error) {
	if len(slc) < 2 {
		return nil, errors.New("Tail called on empty list")
	}
	return slc[1:], nil
}
