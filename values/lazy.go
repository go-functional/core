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
