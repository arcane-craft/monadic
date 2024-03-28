package result

import (
	"fmt"

	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
)

func Ok[A any](v lazy.Value[A]) either.Either[A, error] {
	return either.Left[A, error](v)
}

func Fail[A any](err lazy.Value[error]) either.Either[A, error] {
	return either.Right[A](err)
}

func IsOk[A any](et lazy.Value[either.Either[A, error]]) lazy.Value[bool] {
	return either.IsLeft(et)
}

func IsFail[A any](et lazy.Value[either.Either[A, error]]) lazy.Value[bool] {
	return either.IsRight(et)
}

func FromOk[A any](v lazy.Value[either.Either[A, error]], def lazy.Value[A]) lazy.Value[A] {
	return either.FromLeft(v, def)
}

func FromFail[A any](v lazy.Value[either.Either[A, error]]) lazy.Value[error] {
	return either.FromRight(v, lazy.Const[error](nil))
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

func WrapFail[A any](msg lazy.Value[string], fa lazy.Value[either.Either[A, error]]) lazy.Value[either.Either[A, error]] {
	return MapFail(fa, lazy.Lift(func(err error) error {
		return fmt.Errorf("%s %w", lazy.Eval(msg), err)
	}))
}
