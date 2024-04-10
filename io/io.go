package io

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/monad"
)

type eType struct {
	error
}

func (e eType) IsNil() bool {
	return e.error == nil
}

type rIO[A any, _E monadic.Nillable] struct {
	v lazy.Value[either.Either[error, A]]
}

type IO[A any] rIO[A, eType]

func New[A any](f func() either.Either[error, A]) IO[A] {
	return IO[A]{lazy.New(f)}
}

var _ = monad.ImplMonadDoClass[IO[any]]()
