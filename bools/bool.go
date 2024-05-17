package bools

import "github.com/arcane-craft/monadic/lazy"

type Bool bool

func IfThenElse[A any](b Bool, t lazy.Value[A], e lazy.Value[A]) A {
	if b {
		return t()
	}
	return e()
}

func Or(a Bool, b lazy.Value[Bool]) Bool {
	return a || b()
}

func And(a Bool, b lazy.Value[Bool]) Bool {
	return a && b()
}

func Not(a Bool) Bool {
	return !a
}
