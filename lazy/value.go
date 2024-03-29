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

func Zero[A any]() Value[A] {
	var zero A
	return Const(zero)
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

func If[A any](ok Bool, ret Value[A], def Value[A]) Value[A] {
	return Bind(ok, func(b bool) Value[A] {
		if b {
			return ret
		}
		return def
	})
}

func IfThen[A any](ok Bool, then func() Value[A], def Value[A]) Value[A] {
	return Bind(ok, func(b bool) Value[A] {
		if b {
			return then()
		}
		return def
	})
}

func IfThenElse[A any](ok Bool, then func() Value[A], elsef func() Value[A]) Value[A] {
	return Bind(ok, func(b bool) Value[A] {
		if b {
			return then()
		}
		return elsef()
	})
}

func And(a, b Bool) Bool {
	return New(func() bool {
		return Eval(a) && Eval(b)
	})
}

func Or(a, b Bool) Bool {
	return New(func() bool {
		return Eval(a) || Eval(b)
	})
}

func All(a Bool, left ...Bool) Bool {
	return New(func() bool {
		if !Eval(a) {
			return false
		}
		for _, b := range left {
			if !Eval(b) {
				return false
			}
		}
		return true
	})
}

func Any(a Bool, left ...Bool) Bool {
	return New(func() bool {
		if Eval(a) {
			return true
		}
		for _, b := range left {
			if Eval(b) {
				return true
			}
		}
		return false
	})
}
