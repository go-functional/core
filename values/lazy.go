package values

import "sync"

// Lazy is an interface
type Lazy[T any] struct {
	done bool
	val  T
	err  error
	mut  *sync.Mutex
	f    func() (T, error)
}

func NewLazy[T any](run func() (T, error)) *Lazy[T] {
	return &Lazy[T]{
		done: false,
		mut:  &sync.Mutex{},
		f:    run,
	}
}

func GetLazy[T any](l *Lazy[T]) (T, error) {
	l.mut.Lock()
	defer l.mut.Unlock()
	if l.done {
		return l.val, l.err
	}
	l.val, l.err = l.f()
	l.done = true
	return l.val, l.err
}

func MapLazy[T, U any](l *Lazy[T], f func(T) (U, error)) *Lazy[U] {
	l.mut.Lock()
	defer l.mut.Unlock()
	return NewLazy(func() (U, error) {
		t, err := GetLazy(l)
		var zeroU U
		if err != nil {
			return zeroU, err
		}
		return f(t)
	})
}
