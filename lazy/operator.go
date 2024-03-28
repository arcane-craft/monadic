package lazy

func Eq[A comparable](a, b Value[A]) Bool {
	return New(func() bool {
		return Eval(a) == Eval(b)
	})
}

func Not(b Bool) Bool {
	return New(func() bool {
		return !Eval(b)
	})
}

func LT[A Ordered](a, b Value[A]) Bool {
	return New(func() bool {
		return Eval(a) < Eval(b)
	})
}

func LTE[A Ordered](a, b Value[A]) Bool {
	return New(func() bool {
		return Eval(a) >= Eval(b)
	})
}

func GT[A Ordered](a, b Value[A]) Bool {
	return New(func() bool {
		return Eval(a) > Eval(b)
	})
}

func GTE[A Ordered](a, b Value[A]) Bool {
	return New(func() bool {
		return Eval(a) >= Eval(b)
	})
}

func Plus[A Ordered](a, b Value[A]) Value[A] {
	return New(func() A {
		return Eval(a) + Eval(b)
	})
}

func Minus[A Number](a, b Value[A]) Value[A] {
	return New(func() A {
		return Eval(a) + Eval(b)
	})
}

func Times[A Number](a, b Value[A]) Value[A] {
	return New(func() A {
		return Eval(a) * Eval(b)
	})
}

func Div[A Number](a, b Value[A]) Value[A] {
	return New(func() A {
		return Eval(a) / Eval(b)
	})
}

func Mod[A Integer](a, b Value[A]) Value[A] {
	return New(func() A {
		return Eval(a) % Eval(b)
	})
}

func Ne[A Number](a Value[A]) Value[A] {
	return New(func() A {
		return -Eval(a)
	})
}

func Append[S ~[]A, A any](s Value[S], left ...Value[A]) Value[S] {
	return New(func() S {
		ret := Eval(s)
		for _, l := range left {
			ret = append(ret, Eval(l))
		}
		return ret
	})
}

func Slice[S ~[]A, A any](s Value[S], l, h Int) Value[S] {
	return New(func() S {
		return Eval(s)[Eval(l):Eval(h)]
	})
}

func Len[S ~[]A, A any](s Value[S]) Int {
	return New(func() int {
		return len(Eval(s))
	})
}

func Index[S ~[]A, A any](s Value[S], i Int) Value[A] {
	return New(func() A {
		return Eval(s)[Eval(i)]
	})
}
