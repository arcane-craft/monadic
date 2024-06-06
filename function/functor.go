package function

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/basics"
)

func Map[R, A, B any](m func(A) B, f func(R) A) func(R) B {
	return Compose(m, f)
}

func Replace[R, A, B any](a A, f func(R) B) func(R) A {
	return Map(func(B) A {
		return a
	}, f)
}

func Ignore[R, A any](f func(R) A) func(R) monadic.Unit {
	return Map(Partial(basics.Const[A], monadic.Unit{}), f)
}
