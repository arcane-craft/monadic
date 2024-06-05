package tuple

import "github.com/arcane-craft/monadic/function"

func Match[R, A, B any](f func(A, B) R, t Tuple[A, B]) R {
	return f(t())
}

func Match3[R, A, B, C any](f func(A, B, C) R, t Tuple3[A, B, C]) R {
	return f(t())
}

func Match4[R, A, B, C, D any](f func(A, B, C, D) R, t Tuple4[A, B, C, D]) R {
	return f(t())
}

func Match5[R, A, B, C, D, E any](f func(A, B, C, D, E) R, t Tuple5[A, B, C, D, E]) R {
	return f(t())
}

func MatchF[R, A, B any](f func(A, B) R) func(Tuple[A, B]) R {
	return function.Partial(Match, f)
}

func MatchF3[R, A, B, C any](f func(A, B, C) R) func(Tuple3[A, B, C]) R {
	return function.Partial(Match3, f)
}

func MatchF4[R, A, B, C, D any](f func(A, B, C, D) R) func(Tuple4[A, B, C, D]) R {
	return function.Partial(Match4, f)
}

func MatchF5[R, A, B, C, D, E any](f func(A, B, C, D, E) R) func(Tuple5[A, B, C, D, E]) R {
	return function.Partial(Match5, f)
}
