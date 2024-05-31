package exception

import (
	"fmt"

	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/io"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/monad"
	"github.com/arcane-craft/monadic/result"
	"github.com/arcane-craft/monadic/tuple"
)

func Try[A any](m io.IO[A]) io.IO[either.Either[error, A]] {
	return Catch(
		monad.LiftM[io.IO[either.Either[error, A]]](result.Ok, m),
		function.Compose(applicative.Pure[io.IO[either.Either[error, A]]], result.Fail[A]))
}

func Catch[A any](m io.IO[A], h func(error) io.IO[A]) io.IO[A] {
	return io.CatchIO(m, h)
}

func Descript[A any](m io.IO[A], desc string) io.IO[A] {
	return io.New(lazy.New(func() either.Either[error, A] {
		return result.Fail[A](fmt.Errorf(desc+" failed: %w", tuple.T2nd(io.Perform(m))))
	}))
}

func Throw[A any](err error) io.IO[A] {
	return io.New(func() either.Either[error, A] {
		return result.Fail[A](err)
	})
}

func Bracket[A, B, C any](before io.IO[A], after func(A) io.IO[B], thing func(A) io.IO[C]) io.IO[C] {
	return monad.Bind(before, func(a A) io.IO[C] {
		return monad.Bind(Catch(thing(a), func(err error) io.IO[C] {
			return monad.Then_(after(a), Throw[C](err))
		}), func(c C) io.IO[C] {
			return monad.Then_(after(a), applicative.Pure[io.IO[C]](c))
		})
	})
}

func Bracket_[A, B, C any](before io.IO[A], after io.IO[B], thing io.IO[C]) io.IO[C] {
	return Bracket(before, func(A) io.IO[B] {
		return after
	}, func(A) io.IO[C] {
		return thing
	})
}

func Finally[A, B any](thing io.IO[A], final io.IO[B]) io.IO[A] {
	return monad.Bind(Catch(thing, func(err error) io.IO[A] {
		return monad.Then_(final, Throw[A](err))
	}), func(a A) io.IO[A] {
		return monad.Then_(final, applicative.Pure[io.IO[A]](a))
	})
}
