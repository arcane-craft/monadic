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
	var zero A
	return monad.Bind(applicative.Pure[io.IO[A]](zero), func(A) io.IO[A] {
		ret, err := io.Perform(m)
		if err != nil {
			return h(err)
		}
		return applicative.Pure[io.IO[A]](ret)
	})
}

func Descript[A any](m io.IO[A], desc string) io.IO[A] {
	return io.New(lazy.New(func() either.Either[error, A] {
		return result.Fail[A](fmt.Errorf(desc+" failed: %w", tuple.T2nd(m.Resolve())))
	}))
}

func Throw[A any](err error) io.IO[A] {
	return io.New(func() either.Either[error, A] {
		return result.Fail[A](err)
	})
}
