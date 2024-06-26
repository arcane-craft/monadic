package tuple

type Tuple[A, B any] func() (A, B)

func T[A, B any](a A, b B) Tuple[A, B] {
	return func() (A, B) {
		return a, b
	}
}

func First[A, B any](a A, _ B) A {
	return a
}

func Second[A, B any](_ A, b B) B {
	return b
}

func Map[A, B, Ap, Bp any](m func(A, B) (Ap, Bp), fa Tuple[A, B]) Tuple[Ap, Bp] {
	return func() (Ap, Bp) {
		return m(fa())
	}
}

type Tuple3[A, B, C any] func() (A, B, C)

func T3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return func() (A, B, C) {
		return a, b, c
	}
}

func First3[A, B, C any](a A, _ B, _ C) A {
	return a
}

func Second3[A, B, C any](_ A, b B, _ C) B {
	return b
}

func Third[A, B, C any](_ A, _ B, c C) C {
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

func First4[A, B, C, D any](a A, _ B, _ C, _ D) A {
	return a
}

func Second4[A, B, C, D any](_ A, b B, _ C, _ D) B {
	return b
}

func Third4[A, B, C, D any](_ A, _ B, c C, _ D) C {
	return c
}

func Fourth[A, B, C, D any](_ A, _ B, _ C, d D) D {
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

func First5[A, B, C, D, E any](a A, _ B, _ C, _ D, _ E) A {
	return a
}

func Second5[A, B, C, D, E any](_ A, b B, _ C, _ D, _ E) B {
	return b
}

func Third5[A, B, C, D, E any](_ A, _ B, c C, _ D, _ E) C {
	return c
}

func Fourth5[A, B, C, D, E any](_ A, _ B, _ C, d D, _ E) D {
	return d
}

func Fifth[A, B, C, D, E any](_ A, _ B, _ C, _ D, e E) E {
	return e
}

func Map5[A, B, C, D, E, Ap, Bp, Cp, Dp, Ep any](m func(A, B, C, D, E) (Ap, Bp, Cp, Dp, Ep), fa Tuple5[A, B, C, D, E]) Tuple5[Ap, Bp, Cp, Dp, Ep] {
	return func() (Ap, Bp, Cp, Dp, Ep) {
		return m(fa())
	}
}
