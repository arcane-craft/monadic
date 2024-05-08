package result

import (
	"fmt"

	"github.com/arcane-craft/monadic/either"
)

func Fail[A any](err error) either.Either[error, A] {
	return either.Left[A](err)
}

func Ok[A any](v A) either.Either[error, A] {
	return either.Right[error](v)
}

func IsFail[A any](e either.Either[error, A]) bool {
	return either.IsLeft(e)
}

func IsOk[A any](e either.Either[error, A]) bool {
	return either.IsRight(e)
}

func FromFail[A any](err error, e either.Either[error, A]) error {
	return either.FromLeft(err, e)
}

func FromOk[A any](b A, e either.Either[error, A]) A {
	return either.FromRight(b, e)
}

func MapFail[A any](m func(error) error, fa either.Either[error, A]) either.Either[error, A] {
	return either.MapLeft(m, fa)
}

func WrapFail[A any](msg string, e either.Either[error, A]) either.Either[error, A] {
	return MapFail(func(err error) error {
		return fmt.Errorf("%s %w", msg, err)
	}, e)
}
