package option

import (
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/monad"
)

type aType struct{}

type rOption[A any, _T any] struct {
	v *A
}

type Option[A any] rOption[A, aType]

func Some[A any](a A) Option[A] {
	return Option[A]{
		v: &a,
	}
}

func None[A any]() Option[A] {
	return Option[A]{}
}

func IsNone[A any](o Option[A]) bool {
	return o.v == nil
}

func IsSome[A any](o Option[A]) bool {
	return o.v != nil
}

func FromOption[A any](a lazy.Value[A], o Option[A]) A {
	if o.v == nil {
		return a()
	}
	return *o.v
}

var _ = monad.ImplMonadDoClass[Option[any]]()
