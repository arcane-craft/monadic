package bools

import "github.com/arcane-craft/monadic/lazy"

func IfThenElse[A any](b bool, t lazy.Value[A], e lazy.Value[A]) A {
	if b {
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

func Default[A any](t lazy.Value[A]) func() (lazy.Value[A], lazy.Value[bool]) {
	return Case(lazy.Pure(true), t)
}

func Switch[A any](cases ...func() (lazy.Value[A], lazy.Value[bool])) A {
	for _, c := range cases {
		if t, ok := c(); ok() {
			return t()
		}
	}
	panic("Switch: no condition is satisfied")
}
