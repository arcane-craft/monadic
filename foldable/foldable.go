package foldable

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/algebra"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/bools"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/num"
)

type Foldable[
	F monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	monadic.Generalize[F, A, _T]
	Foldr(func(A, any) any, any, F) any
}

func ImplFoldable[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Unit {
	return monadic.Unit{}
}

func Foldr[
	A, B any,
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	_T any,
](f func(A, B) B, init B, input F) B {
	return basics.Zero[F]().Foldr(func(a A, b any) any {
		return f(a, b.(B))
	}, init, input).(B)
}

func Foldl[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	A, B any,
	_T any,
](f func(B, A) B, init B, input F) B {
	return Foldr(func(a A, b func(B) B) func(B) B {
		return function.Compose(b, function.Partial(function.Flip(f), a))
	}, basics.Id, input)(init)
}

func Concat[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	A algebra.Monoid[A],
	_T any,
](f F) A {
	return Foldr(algebra.Append, algebra.Neutral[A](), f)
}

func ConcatMap[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	M algebra.Monoid[M],
	A any,
	_T any,
](f func(A) M, fa F) M {
	return Foldr(function.Uncurry(function.Compose(function.Curry(algebra.Append[M]), f)), algebra.Neutral[M](), fa)
}

func And[
	F interface {
		Foldable[F, lazy.Value[bools.Bool], _T]
		monadic.Data[lazy.Value[bools.Bool], _T]
	},
	A any,
	_T any,
](f F) bools.Bool {
	return Foldl(bools.And, true, f)
}

func Or[
	F interface {
		Foldable[F, lazy.Value[bools.Bool], _T]
		monadic.Data[lazy.Value[bools.Bool], _T]
	},
	A any,
	_T any,
](f F) bools.Bool {
	return Foldl(bools.Or, false, f)
}

func Any[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](p func(A) bools.Bool, f F) bools.Bool {
	return Foldl(func(x bools.Bool, y A) bools.Bool {
		return x || p(y)
	}, false, f)
}

func All[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](p func(A) bools.Bool, f F) bools.Bool {
	return Foldl(func(x bools.Bool, y A) bools.Bool {
		return x && p(y)
	}, true, f)
}

func Sum[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	A num.Num,
	_T any,
](f F) A {
	init := basics.Zero[A]()
	return Foldr(num.Plus, init, f)
}

func Product[
	F interface {
		Foldable[F, A, _T]
		monadic.Data[A, _T]
	},
	A num.Num,
	_T any,
](f F) A {
	init := basics.Zero[A]()
	init++
	return Foldr(num.Mult, init, f)
}
