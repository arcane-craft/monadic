package applicative

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/foldable"
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
	FA interface {
		Applicative[FA, A, _T]
		monadic.Data[A, _T]
	},
	FB interface {
		Applicative[FB, B, _T]
		monadic.Data[B, _T]
	},
	A, B, C any,
	_T any,
](f func(A, B) C, a FA, b FB) FC {
	return basics.Zero[FC]().Concretize(basics.Zero[FA]().LiftA2(func(a A, b any) any {
		return f(a, b.(B))
	}, a, basics.Zero[FB]().Abstract(b)))
}

type Alternative[
	F monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	Applicative[F, A, _T]
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
