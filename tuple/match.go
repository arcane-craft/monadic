package tuple

func Match[T Tuple[A, B], R, A, B any](f func(A, B) R, t T) R {
	return f(t())
}

func Match3[T Tuple3[A, B, C], R, A, B, C any](f func(A, B, C) R, t T) R {
	return f(t())
}

func Match4[T Tuple4[A, B, C, D], R, A, B, C, D any](f func(A, B, C, D) R, t T) R {
	return f(t())
}

func Match5[T Tuple5[A, B, C, D, E], R, A, B, C, D, E any](f func(A, B, C, D, E) R, t T) R {
	return f(t())
}

func CurryMatch[T Tuple[A, B], R, A, B any](f func(A, B) R) func(T) R {
	return func(t T) R {
		return f(t())
	}
}

func CurryMatch3[T Tuple3[A, B, C], R, A, B, C any](f func(A, B, C) R) func(T) R {
	return func(t T) R {
		return f(t())
	}
}

func CurryMatch4[T Tuple4[A, B, C, D], R, A, B, C, D any](f func(A, B, C, D) R) func(T) R {
	return func(t T) R {
		return f(t())
	}
}

func CurryMatch5[T Tuple5[A, B, C, D, E], R, A, B, C, D, E any](f func(A, B, C, D, E) R) func(T) R {
	return func(t T) R {
		return f(t())
	}
}
