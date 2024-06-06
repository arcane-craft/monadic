package functor

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/function"
)

type Functor[
	F monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	monadic.Generalize[F, A, _T]
	Map(func(A) any, F) monadic.Data[any, _T]
}

func ImplFunctor[
	F interface {
		Functor[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Unit {
	return monadic.Unit{}
}

func Map[
	FB interface {
		Functor[FB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	FA interface {
		Functor[FA, A, _T]
		monadic.Data[A, _T]
	},
	_T any,
](m func(A) B, fa FA) FB {
	return basics.Zero[FB]().Concretize(basics.Zero[FA]().Map(func(a A) any {
		return m(a)
	}, fa))
}

func Replace[
	FA interface {
		Functor[FA, A, _T]
		monadic.Data[A, _T]
	},
	FB interface {
		Functor[FB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](a A, fb FB) FA {
	return Map[FA](func(B) A {
		return a
	}, fb)
}

func Ignore[
	FT interface {
		Functor[FT, monadic.Unit, _T]
		monadic.Data[monadic.Unit, _T]
	},
	FA interface {
		Functor[FA, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](x FA) FT {
	return Map[FT](function.Partial(basics.Const[A], monadic.Unit{}), x)
}
