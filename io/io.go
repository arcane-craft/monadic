package io

import (
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/monad"
	"github.com/arcane-craft/monadic/result"
)

type aType struct{}

type rIO[A any, _T any] struct {
	v lazy.Value[either.Either[error, A]]
}

type IO[A any] rIO[A, aType]

func New[A any](f func() either.Either[error, A]) IO[A] {
	return IO[A]{lazy.New(f)}
}

func CatchIO[A any](m IO[A], h func(error) IO[A]) IO[A] {
	var zero A
	return monad.Bind(applicative.Pure[IO[A]](zero), func(A) IO[A] {
		ret, err := Perform(m)
		if err != nil {
			return h(err)
		}
		return applicative.Pure[IO[A]](ret)
	})
}

func Perform[A any](m IO[A]) (A, error) {
	r := m.v()
	if result.IsOk(r) {
		return result.FromOk(nil, r), nil
	}
	return basics.Zero[A](), result.FromFail(nil, r)
}

var _ = monad.ImplMonadDo[IO[any]]()
