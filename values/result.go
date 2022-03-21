package values

type Result[T any] struct {
	val T
	err error
}

func Ok[T any](val T) *Result[T] {
	return &Result[T]{val: val, err: nil}
}

func Error[T any](err error) *Result[T] {
	return &Result[T]{err: err}
}

func (r *Result[T]) Ok() bool {
	return r.err == nil
}

func MapResult[T, U any](
	r Result[T],
	fn func(T) U,
) *Result[U] {
	if !r.Ok() {
		return Error[U](r.err)
	}
	return Ok(fn(r.val))
}
