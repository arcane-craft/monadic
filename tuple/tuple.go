package tuple

type Tuple[A any] func() A

func T[A any](a A) Tuple[A] {
	return func() A {
		return a
	}
}

func Map[A, Ap any](m func(A) Ap, fa Tuple[A]) Tuple[Ap] {
	return func() Ap {
		return m(fa())
	}
}

type Tuple2[A, B any] func() (A, B)

func T2[A, B any](a A, b B) Tuple2[A, B] {
	return func() (A, B) {
		return a, b
	}
}

func T1st[A, B any](a A, _ B) A {
	return a
}

func T2nd[A, B any](_ A, b B) B {
	return b
}

func Map2[A, B, Ap, Bp any](m func(A, B) (Ap, Bp), fa Tuple2[A, B]) Tuple2[Ap, Bp] {
	return func() (Ap, Bp) {
		return m(fa())
	}
}

type Tuple3[A, B, C any] func() (A, B, C)

func New3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return func() (A, B, C) {
		return a, b, c
	}
}

func T1st3[A, B, C any](a A, _ B, _ C) A {
	return a
}

func T2nd3[A, B, C any](_ A, b B, _ C) B {
	return b
}

func T3rd3[A, B, C any](_ A, _ B, c C) C {
	return c
}

func Map3[A, B, C, Ap, Bp, Cp any](m func(A, B, C) (Ap, Bp, Cp), fa Tuple3[A, B, C]) Tuple3[Ap, Bp, Cp] {
	return func() (Ap, Bp, Cp) {
		return m(fa())
	}
}

type Tuple4[A, B, C, D any] func() (A, B, C, D)

func T4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return func() (A, B, C, D) {
		return a, b, c, d
	}
}

func T1st4[A, B, C, D any](a A, _ B, _ C, _ D) A {
	return a
}

func T2nd4[A, B, C, D any](_ A, b B, _ C, _ D) B {
	return b
}

func T3rd4[A, B, C, D any](_ A, _ B, c C, _ D) C {
	return c
}

func T4th4[A, B, C, D any](_ A, _ B, _ C, d D) D {
	return d
}

func Map4[A, B, C, D, Ap, Bp, Cp, Dp any](m func(A, B, C, D) (Ap, Bp, Cp, Dp), fa Tuple4[A, B, C, D]) Tuple4[Ap, Bp, Cp, Dp] {
	return func() (Ap, Bp, Cp, Dp) {
		return m(fa())
	}
}

type Tuple5[A, B, C, D, E any] func() (A, B, C, D, E)

func T5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return func() (A, B, C, D, E) {
		return a, b, c, d, e
	}
}

func T1st5[A, B, C, D, E any](a A, _ B, _ C, _ D, _ E) A {
	return a
}

func T2nd5[A, B, C, D, E any](_ A, b B, _ C, _ D, _ E) B {
	return b
}

func T3rd5[A, B, C, D, E any](_ A, _ B, c C, _ D, _ E) C {
	return c
}

func T4th5[A, B, C, D, E any](_ A, _ B, _ C, d D, _ E) D {
	return d
}

func T5th5[A, B, C, D, E any](_ A, _ B, _ C, _ D, e E) E {
	return e
}

func Map5[A, B, C, D, E, Ap, Bp, Cp, Dp, Ep any](m func(A, B, C, D, E) (Ap, Bp, Cp, Dp, Ep), fa Tuple5[A, B, C, D, E]) Tuple5[Ap, Bp, Cp, Dp, Ep] {
	return func() (Ap, Bp, Cp, Dp, Ep) {
		return m(fa())
	}
}
