package bools

import "github.com/arcane-craft/monadic/lazy"

type IfExpr[A any] struct {
	p bool
}

type ThenExpr[A any] struct {
	IfExpr[A]
	t lazy.Value[A]
}

func If[A any](p bool) IfExpr[A] {
	return IfExpr[A]{
		p: p,
	}
}

func (m IfExpr[A]) Then(t lazy.Value[A]) ThenExpr[A] {
	return ThenExpr[A]{
		IfExpr: m,
		t:      t,
	}
}

func (m ThenExpr[A]) Else(e lazy.Value[A]) A {
	if m.p {
		return m.t()
	}
	return e()
}

func IfThenElse[A any](p bool, t lazy.Value[A], e lazy.Value[A]) A {
	if p {
		return t()
	}
	return e()
}

func Or(a bool, b lazy.Value[bool]) bool {
	return a || b()
}

func And(a bool, b lazy.Value[bool]) bool {
	return a && b()
}

func Not(a bool) bool {
	return !a
}

func Case[A any](p lazy.Value[bool], t lazy.Value[A]) func() (lazy.Value[A], lazy.Value[bool]) {
	return func() (lazy.Value[A], lazy.Value[bool]) {
		return t, p
	}
}

func SimpleCase[A any](p bool, t lazy.Value[A]) func() (lazy.Value[A], lazy.Value[bool]) {
	return func() (lazy.Value[A], lazy.Value[bool]) {
		return t, lazy.Pure(p)
	}
}

func Default[A any](t lazy.Value[A]) func() (lazy.Value[A], lazy.Value[bool]) {
	return SimpleCase(true, t)
}

func Switch[A any](cases ...func() (lazy.Value[A], lazy.Value[bool])) A {
	for _, c := range cases {
		if t, ok := c(); ok() {
			return t()
		}
	}
	panic("Switch: no condition is satisfied")
}
