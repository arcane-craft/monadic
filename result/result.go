package result

import (
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
)

func Ok[A any](v lazy.Value[A]) lazy.Value[either.Either[A, error]] {
	return either.Left[A, error](v)
}

func Fail[A any](err lazy.Value[error]) lazy.Value[either.Either[A, error]] {
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

func Bind[A, B any](ma lazy.Value[either.Either[A, error]], mm func(lazy.Value[A]) lazy.Value[either.Either[B, error]]) lazy.Value[either.Either[B, error]] {
	return lazy.Bind(IsFail(ma), func(fail bool) lazy.Value[either.Either[B, error]] {
		if fail {
			return Fail[B](FromFail(ma))
		}
		var zero A
		return mm(FromOk(ma, lazy.Const(zero)))
	})
}
