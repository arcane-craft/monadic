package lazy

func Lift[A, B any](f func(A) B) func(Value[A]) Value[B] {
	return func(v Value[A]) Value[B] {
		return New(func() B {
			return f(Eval(v))
		})
	}
}

func Lift2[A, B, C any](f func(A, B) C) func(Value[A], Value[B]) Value[C] {
	return func(a Value[A], b Value[B]) Value[C] {
		return New(func() C {
			return f(Eval(a), Eval(b))
		})
	}
}

func Lift3[A, B, C, D any](f func(A, B, C) D) func(Value[A], Value[B], Value[C]) Value[D] {
	return func(a Value[A], b Value[B], c Value[C]) Value[D] {
		return New(func() D {
			return f(Eval(a), Eval(b), Eval(c))
		})
	}
}

func Lift4[A, B, C, D, E any](f func(A, B, C, D) E) func(Value[A], Value[B], Value[C], Value[D]) Value[E] {
	return func(a Value[A], b Value[B], c Value[C], d Value[D]) Value[E] {
		return New(func() E {
			return f(Eval(a), Eval(b), Eval(c), Eval(d))
		})
	}
}

func Lift5[A, B, C, D, E, F any](f func(A, B, C, D, E) F) func(Value[A], Value[B], Value[C], Value[D], Value[E]) Value[F] {
	return func(a Value[A], b Value[B], c Value[C], d Value[D], e Value[E]) Value[F] {
		return New(func() F {
			return f(Eval(a), Eval(b), Eval(c), Eval(d), Eval(e))
		})
	}
}
