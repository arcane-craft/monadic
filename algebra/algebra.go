package algebra

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/basics"
)

type Semigroup[T any] interface {
	Append(T, T) T
}

func ImplSemigroup[T Semigroup[T]]() monadic.Unit {
	return monadic.Unit{}
}

type Monoid[T any] interface {
	Semigroup[T]
	Neutral() T
}

func ImplMonoid[T Monoid[T]]() monadic.Unit {
	return monadic.Unit{}
}

func Append[T Semigroup[T]](a T, b T) T {
	return basics.Zero[T]().Append(a, b)
}

func Neutral[T Monoid[T]]() T {
	return basics.Zero[T]().Neutral()
}
