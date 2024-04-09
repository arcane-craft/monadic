package lazy

type Any = Value[any]

func ToAny[A any](v Value[A]) Any {
	return New(func() any {
		return any(Eval(v))
	})
}
