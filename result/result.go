package result

import (
	"fmt"

	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
)

func Ok[A any](v lazy.Value[A]) lazy.Value[either.Either[A, error]] {
	return either.Left[A, error](v)
}

func Fail[A any](err lazy.Value[error]) lazy.Value[either.Either[A, error]] {
	return either.Right[A](err)
}

func IsOk[A any](et lazy.Value[either.Either[A, error]]) lazy.Bool {
	return either.IsLeft(et)
}

func IsFail[A any](et lazy.Value[either.Either[A, error]]) lazy.Bool {
	return either.IsRight(et)
}

func FromOk[A any](v lazy.Value[either.Either[A, error]], def lazy.Value[A]) lazy.Value[A] {
	return either.FromLeft(v, def)
}

func FromFail[A any](v lazy.Value[either.Either[A, error]]) lazy.Value[error] {
	return either.FromRight(v, lazy.Zero[error]())
}

func Map[A, B any](fa lazy.Value[either.Either[A, error]], m func(lazy.Value[A]) lazy.Value[B]) lazy.Value[either.Either[B, error]] {
	return either.Map(fa, m)
}

func Bind[A, B any](ma lazy.Value[either.Either[A, error]], mm func(lazy.Value[A]) lazy.Value[either.Either[B, error]]) lazy.Value[either.Either[B, error]] {
	return either.Bind(ma, mm)
}

func MapFail[A any](fa lazy.Value[either.Either[A, error]], m func(lazy.Value[error]) lazy.Value[error]) lazy.Value[either.Either[A, error]] {
	return either.MapRight(fa, m)
}

func BindFail[A any](ma lazy.Value[either.Either[A, error]], mm func(lazy.Value[error]) lazy.Value[either.Either[error, A]]) lazy.Value[either.Either[A, error]] {
	return either.BindRight(ma, mm)
}

func WrapFail[A any](msg lazy.String, fa lazy.Value[either.Either[A, error]]) lazy.Value[either.Either[A, error]] {
	return BindFail(fa, func(err lazy.Value[error]) lazy.Value[either.Either[error, A]] {
		return either.Left[error, A](lazy.Lift2(func(s string, e error) error {
			return fmt.Errorf("%s %w", s, e)
		})(msg, err))
	})
}
