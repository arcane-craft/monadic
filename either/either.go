package either

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/monad"
)

type eType[A any] struct {
	e *A
}

func (e eType[A]) IsNil() bool {
	return e.e == nil
}

type rEither[A, B any, _E monadic.Nillable] struct {
	left  *A
	right *B
}

type Either[A, B any] rEither[A, B, eType[A]]

func Left[B, A any](v A) Either[A, B] {
	return Either[A, B]{
		left: &v,
	}
}

func Right[A, B any](v B) Either[A, B] {
	return Either[A, B]{
		right: &v,
	}
}

func IsLeft[A, B any](e Either[A, B]) bool {
	return e.left != nil
}

func IsRight[A, B any](e Either[A, B]) bool {
	return e.right != nil
}

func FromLeft[A, B any](a A, e Either[A, B]) A {
	if e.left != nil {
		return *e.left
	}
	return a
}

func FromRight[A, B any](b B, e Either[A, B]) B {
	if e.right != nil {
		return *e.right
	}
	return b
}

func EitherOf[A, B, C any](left func(A) C, right func(B) C, e Either[A, B]) C {
	if e.left != nil {
		return left(*e.left)
	}
	return right(*e.right)
}

func Swap[A, B any](e Either[A, B]) Either[B, A] {
	return EitherOf(Right[B, A], Left[A, B], e)
}

func MapLeft[A, B, C any](m func(A) B, fa Either[A, C]) Either[B, C] {
	return EitherOf(function.Compose(Left[C, B], m), Right[B, C], fa)
}

func MapRight[A, B, C any](m func(B) C, fa Either[A, B]) Either[A, C] {
	return EitherOf(Left[C, A], function.Compose(Right[A, C], m), fa)
}

var _ = monad.ImplMonadDoClass[Either[any, any]]()
