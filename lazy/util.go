package lazy

func Lift[A, B any](f func(A) B) func(Value[A]) Value[B] {
	return func(v Value[A]) Value[B] {
		return New(func() B {
			return f(Eval(v))
		})
	}
}
