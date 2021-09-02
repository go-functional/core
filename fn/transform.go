package fn

// Compose takes two functions and returns one function that calls both.
// The returned function calls fn2 with the output of fn1.
//
// Compose is roughly the opposite of Curry. Where Curry "expands" one function
// into two, Compose "compresses" two functions down into one.
//
// Example usage:
//
//	composedFn := Compose(
//		func(t int) string { return strconv.Itoa(t) },
//		func(u string) string { return fmt.Sprintf("str-%s", u) },
//	)
//	answer := composedFn(123)
//	// answer will be "str-123"
func Compose[T, U, V any](fn1 func(T) U, fn2 func(U) V) func(T) V {
	return func(t T) V {
		return fn2(fn1(t))
	}
}

// Curry takes one function with two parameters and returns a single-parameter
// function that in turn returns a second single parameter function, which then
// returns the value of the original function.
//
// Example usage:
//
//	curriedFn := Curry(func(t int, u string) string {
//		return fmt.Sprintf("%d-%s", t, u)
//	}
//	answer := curriedFn(1)("two")
//	// answer will be "1-2"
//
// Compose is the roughly the opposite of Curry. Where Curry "expands" one function
// into twi, Compose "compresses" two functions down into one.
func Curry[T, U, V any](fn func(T, U) V) func(T) func(U) V {
	return func(t T) func(U) V {
		return func(u U) V {
			return fn(t, u)
		}
	}
}
