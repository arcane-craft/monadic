package either

import (
	"github.com/arcane-craft/monadic/lazy"
)

type Either[A, B any] struct {
	left   lazy.Value[A]
	right  lazy.Value[B]
	isLeft bool
}

func Left[A, B any](v lazy.Value[A]) lazy.Value[Either[A, B]] {
	return lazy.Const(
		Either[A, B]{
			left:   v,
			isLeft: true,
		},
	)
}

func Right[A, B any](v lazy.Value[B]) lazy.Value[Either[A, B]] {
	return lazy.Const(
		Either[A, B]{
			right: v,
		},
	)
}

func IsLeft[A, B any](et lazy.Value[Either[A, B]]) lazy.Value[bool] {
	return lazy.Map(et, func(e Either[A, B]) bool {
		return e.isLeft
	})
}

func IsRight[A, B any](et lazy.Value[Either[A, B]]) lazy.Value[bool] {
	return lazy.Map(et, func(e Either[A, B]) bool {
		return !e.isLeft
	})
}

func FromLeft[A, B any](et lazy.Value[Either[A, B]], def lazy.Value[A]) lazy.Value[A] {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Value[A] {
		if e.isLeft {
			return e.left
		}
		return def
	})
}

func FromRight[A, B any](et lazy.Value[Either[A, B]], def lazy.Value[B]) lazy.Value[B] {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Value[B] {
		if !e.isLeft {
			return e.right
		}
		return def
	})
}

func EitherOf[A, B, C any](et lazy.Value[Either[A, B]], left func(lazy.Value[A]) lazy.Value[C], right func(lazy.Value[B]) lazy.Value[C]) lazy.Value[C] {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Value[C] {
		if !e.isLeft {
			return right(e.right)
		}
		return left(e.left)
	})
}

func Bind[A, B, C any](ma lazy.Value[Either[A, B]], mm func(lazy.Value[A]) lazy.Value[Either[C, B]]) lazy.Value[Either[C, B]] {
	return lazy.Bind(ma, func(e Either[A, B]) lazy.Value[Either[C, B]] {
		return mm(e.left)
	})
}
