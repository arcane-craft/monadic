package applicative

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/functor"
)

type Applicative[
	F monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	functor.Functor[F, A, _T]
	Pure(A) F
	LiftA2(func(A, any) any, F, monadic.Data[any, _T]) monadic.Data[any, _T]
}

func ImplApplicative[
	F interface {
		Applicative[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Unit {
	return monadic.Unit{}
}

func Pure[
	F interface {
		Applicative[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](a A) F {
	return basics.Zero[F]().Pure(a)
}

// <*>
func Apply[
	FB interface {
		Applicative[FB, B, _T]
		monadic.Data[B, _T]
	},
	FM interface {
		Applicative[FM, func(A) B, _T]
		monadic.Data[func(A) B, _T]
	},
	FA interface {
		Applicative[FA, A, _T]
		monadic.Data[A, _T]
	},
	A, B any,
	_T any,
](fm FM, fa FA) FB {
	return LiftA2[FB](function.Uncurry(basics.Id[func(A) B]), fm, fa)
}

// <*
func ApplyL[
	FA interface {
		Applicative[FA, A, _T]
		monadic.Data[A, _T]
	},
	FB interface {
		Applicative[FB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](fa FA, _ FB) FA {
	return LiftA[FA](basics.Id, fa)
}

// *>
func ApplyR[
	FA interface {
		Applicative[FA, A, _T]
		monadic.Data[A, _T]
	},
	FB interface {
		Applicative[FB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](_ FA, fb FB) FB {
	return LiftA[FB](basics.Id, fb)
}

func LiftA[
	FB interface {
		Applicative[FB, B, _T]
		monadic.Data[B, _T]
	},
	FA interface {
		Applicative[FA, A, _T]
		monadic.Data[A, _T]
	},
	A, B any,
	_T any,
](f func(A) B, a FA) FB {
	return LiftA2[FB](func(a A, _ A) B {
		return f(a)
	}, a, a)
}

func LiftA2[
	FC interface {
		Applicative[FC, C, _T]
		monadic.Data[C, _T]
	},
	A, B, C any,
	FA interface {
		Applicative[FA, A, _T]
		monadic.Data[A, _T]
	},
	FB interface {
		Applicative[FB, B, _T]
		monadic.Data[B, _T]
	},
	_T any,
](f func(A, B) C, a FA, b FB) FC {
	return basics.Zero[FC]().Concretize(basics.Zero[FA]().LiftA2(func(a A, b any) any {
		return f(a, b.(B))
	}, a, basics.Zero[FB]().Abstract(b)))
}
