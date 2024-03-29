package function

func Apply[A, B any](f func(A) B, a A) B {
	return f(a)
}

func Apply2[A, B, C any](f func(A, B) C, a A, b B) C {
	return f(a, b)
}

func Apply3[A, B, C, D any](f func(A, B, C) D, a A, b B, c C) D {
	return f(a, b, c)
}

func Apply4[A, B, C, D, E any](f func(A, B, C, D) E, a A, b B, c C, d D) E {
	return f(a, b, c, d)
}

func Apply5[A, B, C, D, E, F any](f func(A, B, C, D, E) F, a A, b B, c C, d D, e E) F {
	return f(a, b, c, d, e)
}

func Partial[A, B, C any](f func(A, B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}

func Partial3[A, B, C, D any](f func(A, B, C) D, a A) func(B, C) D {
	return func(b B, c C) D {
		return f(a, b, c)
	}
}

func Partial32[A, B, C, D any](f func(A, B, C) D, a A, b B) func(C) D {
	return func(c C) D {
		return f(a, b, c)
	}
}

func Partial4[A, B, C, D, E any](f func(A, B, C, D) E, a A) func(B, C, D) E {
	return func(b B, c C, d D) E {
		return f(a, b, c, d)
	}
}

func Partial42[A, B, C, D, E any](f func(A, B, C, D) E, a A, b B) func(C, D) E {
	return func(c C, d D) E {
		return f(a, b, c, d)
	}
}

func Partial43[A, B, C, D, E any](f func(A, B, C, D) E, a A, b B, c C) func(D) E {
	return func(d D) E {
		return f(a, b, c, d)
	}
}

func Partial5[A, B, C, D, E, F any](f func(A, B, C, D, E) F, a A) func(B, C, D, E) F {
	return func(b B, c C, d D, e E) F {
		return f(a, b, c, d, e)
	}
}

func Partial52[A, B, C, D, E, F any](f func(A, B, C, D, E) F, a A, b B) func(C, D, E) F {
	return func(c C, d D, e E) F {
		return f(a, b, c, d, e)
	}
}

func Partial53[A, B, C, D, E, F any](f func(A, B, C, D, E) F, a A, b B, c C) func(D, E) F {
	return func(d D, e E) F {
		return f(a, b, c, d, e)
	}
}

func Partial54[A, B, C, D, E, F any](f func(A, B, C, D, E) F, a A, b B, c C, d D) func(E) F {
	return func(e E) F {
		return f(a, b, c, d, e)
	}
}
