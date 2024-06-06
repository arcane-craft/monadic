package io

import (
	"fmt"

	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/functor"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/monad"
	"github.com/arcane-craft/monadic/result"
)

func (IO[A]) Kind() aType {
	return aType{}
}

func (IO[A]) Concretize(o monadic.Data[any, aType]) IO[A] {
	oi := o.(IO[any])
	return New(func() either.Either[error, A] {
		return basics.Zero[either.Either[error, A]]().Concretize(oi.v())
	})
}

func (IO[A]) Abstract(o IO[A]) monadic.Data[any, aType] {
	return New(func() either.Either[error, any] {
		return basics.Zero[either.Either[error, A]]().Abstract(o.v()).(either.Either[error, any])
	})
}

func (m IO[A]) Delay() lazy.Value[IO[A]] {
	panic(monadic.NotSupportForTest)
}

func (IO[A]) Map(m func(A) any, fa IO[A]) monadic.Data[any, aType] {
	return New(func() either.Either[error, any] {
		return functor.Map[either.Either[error, any]](m, fa.v())
	})
}

func (IO[A]) Pure(a A) IO[A] {
	return New(func() either.Either[error, A] {
		return result.Ok(a)
	})
}

func (IO[A]) LiftA2(f func(A, any) any, a IO[A], b monadic.Data[any, aType]) monadic.Data[any, aType] {
	return New(func() either.Either[error, any] {
		eb := b.(IO[any])
		return applicative.LiftA2[either.Either[error, any]](f, a.v(), eb.v())
	})
}

func (IO[A]) Empty() IO[A] {
	return New(func() either.Either[error, A] {
		return result.Fail[A](fmt.Errorf("mzero"))
	})
}

func (IO[A]) Or(a IO[A], b IO[A]) IO[A] {
	return CatchIO(a, func(_ error) IO[A] {
		return b
	})
}

func (IO[A]) Bind(ma IO[A], mm func(A) monadic.Data[any, aType]) monadic.Data[any, aType] {
	return New(func() either.Either[error, any] {
		return monad.Bind(ma.v(), func(a A) either.Either[error, any] {
			return mm(a).(IO[any]).v()
		})
	})
}

func (IO[A]) Do(proc func() IO[A]) IO[A] {
	return New(func() either.Either[error, A] {
		return basics.Zero[either.Either[error, A]]().Do(func() either.Either[error, A] {
			return proc().v()
		})
	})
}

func (m IO[A]) X() A {
	return m.v().X()
}
