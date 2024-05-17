package traversable

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/foldable"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/functor"
)

type Traversable[
	T monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	functor.Functor[T, A, _T]
	foldable.Foldable[T, A, _T]
	Traverse()
}

func ImplTraversable[
	T interface {
		Traversable[T, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Unit {
	return monadic.Unit{}
}

func Traverse_[
	FN interface {
		applicative.Applicative[FN, monadic.Unit, _F]
		monadic.Data[monadic.Unit, _F]
	},
	A any,
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[FB, _F]
	},
	TA interface {
		Traversable[TA, A, _T]
		monadic.Data[A, _T]
	},
	B any,
	_T any,
	_F any,
](f func(A) FB, t TA) FN {
	return foldable.Foldr(function.Uncurry(function.Compose(function.Curry(applicative.ApplyR[FB, FN]), f)), applicative.Pure[FN](monadic.Unit{}), t)
}

func Sequence_[
	FN interface {
		applicative.Applicative[FN, monadic.Unit, _F]
		monadic.Data[monadic.Unit, _F]
	},
	TFA interface {
		Traversable[TFA, FA, _T]
		monadic.Data[FA, _T]
	},
	FA interface {
		applicative.Applicative[FA, A, _F]
		monadic.Data[A, _F]
	},
	A any,
	_T any,
	_F any,
](t TFA) FN {
	return foldable.Foldr(applicative.ApplyR, applicative.Pure[FN](monadic.Unit{}), t)
}

func For_[
	FN interface {
		applicative.Applicative[FN, monadic.Unit, _F]
		monadic.Data[monadic.Unit, _F]
	},
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[FB, _F]
	},
	TA interface {
		Traversable[TA, A, _T]
		monadic.Data[A, _T]
	},
	A, B any,
	_T any,
	_F any,
](t TA, f func(A) FB) FN {
	return Traverse_[FN](f, t)
}

func Traverse[
	FTB interface {
		applicative.Applicative[FTB, TB, _F]
		monadic.Data[TB, _F]
	},
	A any,
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[B, _F]
	},
	TA interface {
		Traversable[TA, A, _T]
		monadic.Data[A, _T]
	},
	TB interface {
		Traversable[TB, B, _T]
		monadic.Data[B, _T]
	},
	B any,
	_T any,
	_F any,
](f func(A) FB, t TA) FTB {
	panic(monadic.NotSupportForTest)
}

func Sequence[
	FTA interface {
		applicative.Applicative[FTA, TA, _F]
		monadic.Data[TA, _F]
	},
	TFA interface {
		Traversable[TFA, FA, _T]
		monadic.Data[FA, _T]
	},
	TA interface {
		Traversable[TA, A, _T]
		monadic.Data[A, _T]
	},
	FA interface {
		applicative.Applicative[FA, A, _F]
		monadic.Data[A, _F]
	},
	A any,
	_T any,
	_F any,
](t TFA) FTA {
	return Traverse[FTA](basics.Id, t)
}

func For[
	FTB interface {
		applicative.Applicative[FTB, TB, _F]
		monadic.Data[TB, _F]
	},
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[B, _F]
	},
	TB interface {
		Traversable[TB, B, _T]
		monadic.Data[B, _T]
	},
	TA interface {
		Traversable[TA, A, _T]
		monadic.Data[A, _T]
	},
	A, B any,
	_T any,
	_F any,
](t TA, f func(A) FB) FTB {
	return function.Flip(Traverse[FTB, A, FB, TA])(t, f)
}
