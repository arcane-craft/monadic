package io

import (
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/monad"
	"github.com/arcane-craft/monadic/result"
)

func (m IO[A]) Resolve() (A, eType) {
	var zero A
	var e eType
	if m.v == nil {
		return zero, e
	}
	v := m.v()
	if result.IsFail(v) {
		e.error = result.FromFail(nil, v)
		return zero, e
	}
	return result.FromOk(zero, v), e
}

func (IO[A]) Throw(e eType) IO[A] {
	return New(func() either.Either[error, A] {
		return result.Fail[A](e.error)
	})
}

func (IO[A]) Init(f func() IO[A]) IO[A] {
	return New(func() either.Either[error, A] {
		return f().v()
	})
}

func (IO[A]) Pure(a A) IO[A] {
	return New(func() either.Either[error, A] {
		return result.Ok(a)
	})
}

func (m IO[A]) X() A {
	return monad.X(m)
}
