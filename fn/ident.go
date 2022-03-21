package fn

// ID is a function that takes in a value and immediately
// returns it without modification. This functionality is
// useless in most situations, but there are some cases
// in which it's useful.
//
// You'll know them when you see them.
func Identity[T any](t T) T {
	return t
}
