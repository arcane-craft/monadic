package either

import (
	"github.com/arcane-craft/monadic/lazy"
)

type Either[A, B any] struct {
	left   lazy.Value[A]
	right  lazy.Value[B]
	isLeft lazy.Bool
}

func Left[A, B any](v lazy.Value[A]) lazy.Value[Either[A, B]] {
	return lazy.Const(Either[A, B]{
		left:   v,
		isLeft: lazy.Const(true),
	})
}

func Right[A, B any](v lazy.Value[B]) lazy.Value[Either[A, B]] {
	return lazy.Const(Either[A, B]{
		right:  v,
		isLeft: lazy.Const(false),
	})
}

func IsLeft[A, B any](et lazy.Value[Either[A, B]]) lazy.Bool {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Bool {
		return e.isLeft
	})
}

func IsRight[A, B any](et lazy.Value[Either[A, B]]) lazy.Bool {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Bool {
		return lazy.Not(e.isLeft)
	})
}

func FromLeft[A, B any](et lazy.Value[Either[A, B]], def lazy.Value[A]) lazy.Value[A] {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Value[A] {
		return lazy.If(e.isLeft, e.left, def)
	})
}

func FromRight[A, B any](et lazy.Value[Either[A, B]], def lazy.Value[B]) lazy.Value[B] {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Value[B] {
		return lazy.If(lazy.Not(e.isLeft), e.right, def)
	})
}

func EitherOf[A, B, C any](et lazy.Value[Either[A, B]], left func(lazy.Value[A]) lazy.Value[C], right func(lazy.Value[B]) lazy.Value[C]) lazy.Value[C] {
	return lazy.Bind(et, func(e Either[A, B]) lazy.Value[C] {
		return lazy.IfThenElse(e.isLeft,
			func() lazy.Value[C] {
				return left(e.left)
			},
			func() lazy.Value[C] {
				return right(e.right)
			},
		)
	})
}

func Map[A, B, C any](fa lazy.Value[Either[A, B]], m func(lazy.Value[A]) lazy.Value[C]) lazy.Value[Either[C, B]] {
	return lazy.Bind(fa, func(e Either[A, B]) lazy.Value[Either[C, B]] {
		return lazy.IfThenElse(e.isLeft,
			func() lazy.Value[Either[C, B]] {
				return Left[C, B](m(e.left))
			},
			func() lazy.Value[Either[C, B]] {
				return Right[C](e.right)
			},
		)
	})
}

func Bind[A, B, C any](ma lazy.Value[Either[A, B]], mm func(lazy.Value[A]) lazy.Value[Either[C, B]]) lazy.Value[Either[C, B]] {
	return lazy.Bind(ma, func(e Either[A, B]) lazy.Value[Either[C, B]] {
		return lazy.IfThenElse(e.isLeft,
			func() lazy.Value[Either[C, B]] {
				return mm(e.left)
			},
			func() lazy.Value[Either[C, B]] {
				return Right[C](e.right)
			},
		)
	})
}

func Swap[A, B any](v lazy.Value[Either[A, B]]) lazy.Value[Either[B, A]] {
	return lazy.Bind(v, func(e Either[A, B]) lazy.Value[Either[B, A]] {
		return lazy.IfThenElse(e.isLeft,
			func() lazy.Value[Either[B, A]] {
				return Right[B](e.left)
			},
			func() lazy.Value[Either[B, A]] {
				return Left[B, A](e.right)
			},
		)
	})
}

func MapRight[A, B, C any](fa lazy.Value[Either[A, B]], m func(lazy.Value[B]) lazy.Value[C]) lazy.Value[Either[A, C]] {
	return Swap(Map(Swap(fa), func(v lazy.Value[B]) lazy.Value[C] {
		return m(v)
	}))
}

func BindRight[A, B, C any](ma lazy.Value[Either[A, B]], mm func(lazy.Value[B]) lazy.Value[Either[C, A]]) lazy.Value[Either[A, C]] {
	return Swap(Bind(Swap(ma), func(v lazy.Value[B]) lazy.Value[Either[C, A]] {
		return mm(v)
	}))
}
