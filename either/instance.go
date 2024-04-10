package either

import (
	"github.com/arcane-craft/monadic/monad"
)

func (m Either[A, B]) Resolve() (B, eType[A]) {
	var zeroA A
	var zeroB B
	var e eType[A]
	if m.left == nil && m.right == nil {
		return zeroB, e
	}
	if IsLeft(m) {
		left := FromLeft(zeroA, m)
		e.e = &left
		return zeroB, e
	}
	return FromRight(zeroB, m), e
}

func (Either[A, B]) Throw(e eType[A]) Either[A, B] {
	return Left[B](*e.e)
}

func (Either[A, B]) Init(f func() Either[A, B]) Either[A, B] {
	return f()
}

func (Either[A, B]) Pure(b B) Either[A, B] {
	return Right[A](b)
}

func (m Either[A, B]) X() B {
	return monad.X(m)
}
