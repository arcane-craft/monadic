package lazy

type Bool = Value[bool]

var (
	True  = Const(true)
	False = Const(false)
)

func ToBool[B, A ~bool](v Value[A]) Value[B] {
	return New(func() B {
		return B(Eval(v))
	})
}
