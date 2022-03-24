package values

import "sync"

// Lazy is a cached value that will not be computed
// until it is explictly requested by calling GetLazy.
// Thereafter, it is cached and returned from the cached value.
// This construct is best used to hold the result of
// an expensive operation, the return value of which
// doesn't change.
type Lazy[T any] struct {
	done bool
	val  T
	err  error
	mut  *sync.Mutex
	f    func() (T, error)
}

// NewLazy creates a new Lazy value with the given function
// representing the computation to run to get the value
// at some point in the future, when it's requested.
func NewLazy[T any](run func() (T, error)) *Lazy[T] {
	return &Lazy[T]{
		done: false,
		mut:  &sync.Mutex{},
		f:    run,
	}
}

// Get returns the value. If it has already been computed,
// returns it directly from the cache. Otherwise, runs the
// (potentially expensive and long running) computation,
// stores the result in the cache, and returns it.
func (l *Lazy[T]) Get() (T, error) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if l.done {
		return l.val, l.err
	}
	l.val, l.err = l.f()
	l.done = true
	return l.val, l.err
}

// MapLazy returns a new Lazy value that, when requested, will be the
// result of running f on the return value of l.Get.
func MapLazy[T, U any](l *Lazy[T], f func(T) (U, error)) *Lazy[U] {
	l.mut.Lock()
	defer l.mut.Unlock()
	return NewLazy(func() (U, error) {
		t, err := l.Get()
		var zeroU U
		if err != nil {
			return zeroU, err
		}
		return f(t)
	})
}
