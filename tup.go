package core

// Tuple is a structure that holds exactly two values.
// You can read the values, but not change them
type Tuple[T any, U any] struct {
	first T
	second U
}

// Tup creates a new tuple with the first parameter being the first
// element in the tuple and the second being the second
func Tup[T, U any](first T, second U) Tuple[T, U] {
	return Tuple{first: first, second: second}
}

// First gets the first element of the tuple
func First[T any, U any](t Tuple[T, U]) T {
	return t.first
}

// Second gets the second element of the given tuple
func Second[T any, U any](t Tuple[T, U]) U {
	return t.second
}
