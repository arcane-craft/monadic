package either

import (
	"github.com/arcane-craft/monadic/bools"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/monad"
)

type aType[A any] struct {
	e A
}

type rEither[A, B any, _T any] struct {
	left  *A
	right *B
}

type Either[A, B any] rEither[A, B, aType[A]]

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

func IsLeft[A, B any](e Either[A, B]) bools.Bool {
	return e.left != nil
}

func IsRight[A, B any](e Either[A, B]) bools.Bool {
	return e.right != nil
}

func FromLeft[A, B any](a lazy.Value[A], e Either[A, B]) A {
	if e.left != nil {
		return *e.left
	}
	return a()
}

func FromRight[A, B any](b lazy.Value[B], e Either[A, B]) B {
	if e.right != nil {
		return *e.right
	}
	return b()
}

func EitherOf[A, B, C any](left func(A) C, right func(B) C, e Either[A, B]) C {
	if e.left != nil {
		return left(*e.left)
	}
	return right(*e.right)
}

func Mirror[A, B any](e Either[A, B]) Either[B, A] {
	return EitherOf(Right[B, A], Left[A, B], e)
}

func MapLeft[A, B, C any](m func(A) B, fa Either[A, C]) Either[B, C] {
	return EitherOf(function.Compose(Left[C, B], m), Right[B, C], fa)
}

func MapRight[A, B, C any](m func(B) C, fa Either[A, B]) Either[A, C] {
	return EitherOf(Left[C, A], function.Compose(Right[A, C], m), fa)
}

var _ = monad.ImplMonadDo[Either[any, any]]()
