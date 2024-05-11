package functor

import "github.com/arcane-craft/monadic"

type FunctorClass[
	F monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	Map(func(A) any, F) monadic.Data[any, _T]
}

func ImplFunctorClass[
	F interface {
		FunctorClass[F, A, _T]
		monadic.Generalize[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Void {
	return monadic.Void{}
}

func Map[
	FB interface {
		FunctorClass[FB, B, _T]
		monadic.Generalize[FB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	FA interface {
		FunctorClass[FA, A, _T]
		monadic.Generalize[FA, A, _T]
		monadic.Data[A, _T]
	},
	_T any,
](m func(A) B, fa FA) FB {
	return monadic.Zero[FB]().Concretize(monadic.Zero[FA]().Map(func(a A) any {
		return m(a)
	}, fa))
}

func Replace[
	FA interface {
		FunctorClass[FA, A, _T]
		monadic.Generalize[FA, A, _T]
		monadic.Data[A, _T]
	},
	FB interface {
		FunctorClass[FB, B, _T]
		monadic.Generalize[FB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](a A, fb FB) FA {
	return Map[FA](func(B) A {
		return a
	}, fb)
}
