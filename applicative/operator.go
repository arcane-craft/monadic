package applicative

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/functor"
)

type ApplicativeClass[
	F monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	Pure(A) F
	Apply(monadic.Data[func(A) any, _T], F) monadic.Data[any, _T]
}

func ImplApplicativeClass[
	F interface {
		ApplicativeClass[F, A, _T]
		functor.FunctorClass[F, A, _T]
		monadic.Generalize[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Void {
	return monadic.Void{}
}

type AlternativeClass[
	FA monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	Empty() FA
	Or(FA, FA) FA
}

func ImplAlternativeClass[
	F interface {
		AlternativeClass[F, A, _T]
		ApplicativeClass[F, A, _T]
		functor.FunctorClass[F, A, _T]
		monadic.Generalize[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Void {
	return monadic.Void{}
}

func Pure[
	F interface {
		ApplicativeClass[F, A, _T]
		functor.FunctorClass[F, A, _T]
		monadic.Generalize[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](a A) F {
	return monadic.Zero[F]().Pure(a)
}

func Apply[
	FB interface {
		ApplicativeClass[FB, B, _T]
		functor.FunctorClass[FB, B, _T]
		monadic.Generalize[FB, B, _T]
		monadic.Data[B, _T]
	},
	FM interface {
		ApplicativeClass[FM, func(A) B, _T]
		functor.FunctorClass[FM, func(A) B, _T]
		monadic.Generalize[FM, func(A) B, _T]
		monadic.Data[func(A) B, _T]
	},
	FA interface {
		ApplicativeClass[FA, A, _T]
		functor.FunctorClass[FA, A, _T]
		monadic.Generalize[FA, A, _T]
		monadic.Data[A, _T]
	},
	A, B any,
	_T any,
](fm FM, fa FA) FB {
	return monadic.Zero[FB]().Concretize(monadic.Zero[FA]().Apply(fm, fa))
}

func Empty[
	F interface {
		AlternativeClass[F, A, _T]
		ApplicativeClass[F, A, _T]
		functor.FunctorClass[F, A, _T]
		monadic.Generalize[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() F {
	return monadic.Zero[F]().Empty()
}

func Or[
	F interface {
		AlternativeClass[F, A, _T]
		ApplicativeClass[F, A, _T]
		functor.FunctorClass[F, A, _T]
		monadic.Generalize[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](a F, b F) F {
	return monadic.Zero[F]().Or(a, b)
}
