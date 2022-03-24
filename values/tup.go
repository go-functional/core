package values

// Tuple is a structure that holds exactly two values.
// You can read the values, but not change them
type Tuple[T, U any] struct {
	first  T
	second U
}

// Tup creates a new tuple with the first parameter being the first
// element in the tuple and the second being the second
func NewTuple[T, U any](first T, second U) Tuple[T, U] {
	return Tuple[T, U]{
		first:  first,
		second: second,
	}
}

// First gets the first element of the tuple
func (tup *Tuple[T, U]) GetFirst() T {
	return tup.first
}

func (tup *Tuple[T, U]) SetFirst(val T) {
	tup.first = val
}

func (tup *Tuple[T, U]) SetSecond(val U) {
	tup.second = val
}

// Second gets the second element of the given tuple
func (tup *Tuple[T, U]) GetSecond() U {
	return tup.second
}
