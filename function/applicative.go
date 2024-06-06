package function

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/lazy"
)

func Pure[R, A any](a A) func(R) A {
	return Partial(basics.Const[R], a)
}

func FApply[R, A, B any](f func(R) func(A) B, a func(R) A) func(R) B {
	return func(r R) B {
		return f(r)(a(r))
	}
}

func ApplyL[R, A, B any](a func(R) A, _ func(R) B) func(R) A {
	return LiftA(basics.Id, a)
}

func ApplyR[R, A, B any](_ func(R) A, b func(R) B) func(R) B {
	return LiftA(basics.Id, b)
}

func LiftA[R, A, B any](f func(A) B, a func(R) A) func(R) B {
	return func(r R) B {
		return f(a(r))
	}
}

func LiftA2[R, A, B, C any](f func(A, B) C, a func(R) A, b func(R) B) func(R) C {
	return func(r R) C {
		return f(a(r), b(r))
	}
}

func When[R, A any](p bool, f lazy.Value[func(R) monadic.Unit]) func(R) monadic.Unit {
	if p {
		return f()
	}
	return Pure[R](monadic.Unit{})
}
