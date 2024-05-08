package option

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/monad"
)

type eType bool

func (e eType) IsNil() bool {
	return !bool(e)
}

type rOption[A any, _E monadic.Nillable] struct {
	v *A
}

type Option[A any] rOption[A, eType]

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

var _ = monad.ImplMonadDoClass[Option[any]]()
