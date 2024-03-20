package lazy

type Value[A any] struct {
	f func() A
}

func New[A any](f func() A) Value[A] {
	var ret *A
	return Value[A]{
		f: func() A {
			if ret == nil {
				ret = new(A)
				*ret = f()
			}
			return *ret
		},
	}
}

func Const[A any](v A) Value[A] {
	return Value[A]{
		f: func() A {
			return v
		},
	}
}

func Eval[A any](v Value[A]) A {
	return v.f()
}

func Map[A, B any](fa Value[A], m func(A) B) Value[B] {
	return New(func() B {
		return m(Eval(fa))
	})
}

func Bind[A, B any](ma Value[A], mm func(A) Value[B]) Value[B] {
	return New(func() B {
		return Eval(mm(Eval(ma)))
	})
}
