package alternative

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/foldable"
	"github.com/arcane-craft/monadic/function"
)

type Alternative[
	F monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	applicative.Applicative[F, A, _T]
	Empty() F
	Or(F, F) F
}

func ImplAlternative[
	F interface {
		Alternative[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Unit {
	return monadic.Unit{}
}

func Empty[
	F interface {
		Alternative[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() F {
	return basics.Zero[F]().Empty()
}

func Or[
	F interface {
		Alternative[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](a F, b F) F {
	return basics.Zero[F]().Or(a, b)
}

func Choice[
	F interface {
		Alternative[F, A, _T]
		monadic.Data[A, _T]
	},
	T interface {
		foldable.Foldable[T, F, _T]
		monadic.Data[F, _T]
	},
	A any,
	_T any,
](x T) F {
	return foldable.Foldr(Or, Empty[F](), x)
}

func ChoiceMap[
	A any,
	FB interface {
		Alternative[FB, B, _T]
		monadic.Data[B, _T]
	},
	TA interface {
		foldable.Foldable[TA, A, _T]
		monadic.Data[A, _T]
	},
	B any,
	_T any,
](f func(A) FB, x TA) FB {
	return foldable.Foldr(func(elt A, acc FB) FB {
		return function.Infix(f(elt), Or, acc)
	}, Empty[FB](), x)
}
