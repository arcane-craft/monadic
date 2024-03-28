package either

import (
	"github.com/arcane-craft/monadic/lazy"
)

type Either[A, B any] struct {
	left   lazy.Value[A]
	right  lazy.Value[B]
	isLeft lazy.Value[bool]
}

func Left[A, B any](v lazy.Value[A]) Either[A, B] {
	return Either[A, B]{
		left:   v,
		isLeft: lazy.Const(true),
	}
}

func Right[A, B any](v lazy.Value[B]) Either[A, B] {
	return Either[A, B]{
		right:  v,
		isLeft: lazy.Const(false),
	}
}

func IsLeft[A, B any](et lazy.Value[Either[A, B]]) lazy.Value[bool] {
	return lazy.Map(et, func(e Either[A, B]) bool {
		return lazy.Eval(e.isLeft)
	})
}

func IsRight[A, B any](et lazy.Value[Either[A, B]]) lazy.Value[bool] {
	return lazy.Map(et, func(e Either[A, B]) bool {
		return !lazy.Eval(e.isLeft)
	})
}

func FromLeft[A, B any](et lazy.Value[Either[A, B]], def lazy.Value[A]) lazy.Value[A] {
	return lazy.Map(et, func(e Either[A, B]) A {
		if lazy.Eval(e.isLeft) {
			return lazy.Eval(e.left)
		}
		return lazy.Eval(def)
	})
}

func FromRight[A, B any](et lazy.Value[Either[A, B]], def lazy.Value[B]) lazy.Value[B] {
	return lazy.Map(et, func(e Either[A, B]) B {
		if !lazy.Eval(e.isLeft) {
			return lazy.Eval(e.right)
		}
		return lazy.Eval(def)
	})
}

func EitherOf[A, B, C any](et lazy.Value[Either[A, B]], left func(lazy.Value[A]) lazy.Value[C], right func(lazy.Value[B]) lazy.Value[C]) lazy.Value[C] {
	return lazy.Map(et, func(e Either[A, B]) C {
		if !lazy.Eval(e.isLeft) {
			return lazy.Eval(right(e.right))
		}
		return lazy.Eval(left(e.left))
	})
}

func Map[A, B, C any](fa lazy.Value[Either[A, B]], m func(lazy.Value[A]) lazy.Value[C]) lazy.Value[Either[C, B]] {
	return lazy.Map(fa, func(e Either[A, B]) Either[C, B] {
		if !lazy.Eval(e.isLeft) {
			return Right[C](e.right)
		}
		return Left[C, B](m(e.left))
	})
}

func Bind[A, B, C any](ma lazy.Value[Either[A, B]], mm func(lazy.Value[A]) lazy.Value[Either[C, B]]) lazy.Value[Either[C, B]] {
	return lazy.Bind(ma, func(e Either[A, B]) lazy.Value[Either[C, B]] {
		if !lazy.Eval(e.isLeft) {
			return lazy.Const(Right[C](e.right))
		}
		return mm(e.left)
	})
}

func Swap[A, B any](v lazy.Value[Either[A, B]]) lazy.Value[Either[B, A]] {
	return lazy.Map(v, func(e Either[A, B]) Either[B, A] {
		if lazy.Eval(e.isLeft) {
			return Right[B](e.left)
		}
		return Left[B, A](e.right)
	})
}

func MapRight[A, B, C any](fa lazy.Value[Either[A, B]], m func(lazy.Value[B]) lazy.Value[C]) lazy.Value[Either[A, C]] {
	return Swap(Map(Swap(fa), func(v lazy.Value[B]) lazy.Value[C] {
		return m(v)
	}))
}
