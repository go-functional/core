package values

type Option[T any] struct {
	val T
	ok  bool
}

func (o Option[T]) Ok() bool {
	return o.ok
}

func Some[T any](val T) Option[T] {
	return Option[T]{val: val, ok: true}
}

func None[T any]() Option[T] {
	return Option[T]{ok: false}
}

func MapOption[T, U any](
	opt Option[T],
	fn func(T) U,
) Option[U] {
	if !opt.Ok() {
		return None[U]()
	}
	return Some(fn(opt.val))
}
