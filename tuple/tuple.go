package tuple

type Tuple[A any] func() A

func New[A any](a A) Tuple[A] {
	return func() A {
		return a
	}
}

func Map[A, Ap any](fa Tuple[A], m func(A) Ap) Tuple[Ap] {
	return func() Ap {
		return m(fa())
	}
}

type Tuple2[A, B any] func() (A, B)

func New2[A, B any](a A, b B) Tuple2[A, B] {
	return func() (A, B) {
		return a, b
	}
}

func Map2[A, B, Ap, Bp any](fa Tuple2[A, B], m func(A, B) (Ap, Bp)) Tuple2[Ap, Bp] {
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

func Map3[A, B, C, Ap, Bp, Cp any](fa Tuple3[A, B, C], m func(A, B, C) (Ap, Bp, Cp)) Tuple3[Ap, Bp, Cp] {
	return func() (Ap, Bp, Cp) {
		return m(fa())
	}
}

type Tuple4[A, B, C, D any] func() (A, B, C, D)

func New4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return func() (A, B, C, D) {
		return a, b, c, d
	}
}

func Map4[A, B, C, D, Ap, Bp, Cp, Dp any](fa Tuple4[A, B, C, D], m func(A, B, C, D) (Ap, Bp, Cp, Dp)) Tuple4[Ap, Bp, Cp, Dp] {
	return func() (Ap, Bp, Cp, Dp) {
		return m(fa())
	}
}

type Tuple5[A, B, C, D, E any] func() (A, B, C, D, E)

func New5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return func() (A, B, C, D, E) {
		return a, b, c, d, e
	}
}

func Map5[A, B, C, D, E, Ap, Bp, Cp, Dp, Ep any](fa Tuple5[A, B, C, D, E], m func(A, B, C, D, E) (Ap, Bp, Cp, Dp, Ep)) Tuple5[Ap, Bp, Cp, Dp, Ep] {
	return func() (Ap, Bp, Cp, Dp, Ep) {
		return m(fa())
	}
}

type Tuple6[A, B, C, D, E, F any] func() (A, B, C, D, E, F)

func New6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return func() (A, B, C, D, E, F) {
		return a, b, c, d, e, f
	}
}

func Map6[A, B, C, D, E, F, Ap, Bp, Cp, Dp, Ep, Fp any](fa Tuple6[A, B, C, D, E, F], m func(A, B, C, D, E, F) (Ap, Bp, Cp, Dp, Ep, Fp)) Tuple6[Ap, Bp, Cp, Dp, Ep, Fp] {
	return func() (Ap, Bp, Cp, Dp, Ep, Fp) {
		return m(fa())
	}
}

type Tuple7[A, B, C, D, E, F, G any] func() (A, B, C, D, E, F, G)

func New7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return func() (A, B, C, D, E, F, G) {
		return a, b, c, d, e, f, g
	}
}

func Map7[A, B, C, D, E, F, G, Ap, Bp, Cp, Dp, Ep, Fp, Gp any](fa Tuple7[A, B, C, D, E, F, G], m func(A, B, C, D, E, F, G) (Ap, Bp, Cp, Dp, Ep, Fp, Gp)) Tuple7[Ap, Bp, Cp, Dp, Ep, Fp, Gp] {
	return func() (Ap, Bp, Cp, Dp, Ep, Fp, Gp) {
		return m(fa())
	}
}

type Tuple8[A, B, C, D, E, F, G, H any] func() (A, B, C, D, E, F, G, H)

func New8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return func() (A, B, C, D, E, F, G, H) {
		return a, b, c, d, e, f, g, h
	}
}

func Map8[A, B, C, D, E, F, G, H, Ap, Bp, Cp, Dp, Ep, Fp, Gp, Hp any](fa Tuple8[A, B, C, D, E, F, G, H], m func(A, B, C, D, E, F, G, H) (Ap, Bp, Cp, Dp, Ep, Fp, Gp, Hp)) Tuple8[Ap, Bp, Cp, Dp, Ep, Fp, Gp, Hp] {
	return func() (Ap, Bp, Cp, Dp, Ep, Fp, Gp, Hp) {
		return m(fa())
	}
}
