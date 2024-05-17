package either

import (
	"github.com/arcane-craft/monadic"
)

func (Either[A, B]) Kind() aType[A] {
	return aType[A]{}
}

func (Either[A, B]) Concretize(o monadic.Data[any, aType[A]]) Either[A, B] {
	oi := o.(Either[A, any])
	if IsRight(oi) {
		return Right[A]((*oi.right).(B))
	}
	return Left[B](*oi.left)
}

func (Either[A, B]) Abstract(o Either[A, B]) monadic.Data[any, aType[A]] {
	if IsRight(o) {
		return Right[A](any(*o.right))
	}
	return Left[any](*o.left)
}

func (Either[A, B]) Map(m func(B) any, fa Either[A, B]) monadic.Data[any, aType[A]] {
	if IsLeft(fa) {
		return Left[any](*fa.left)
	}
	return Right[A](m(*fa.right))
}

func (Either[A, B]) Pure(b B) Either[A, B] {
	return Right[A](b)
}

func (Either[A, B]) LiftA2(f func(B, any) any, a Either[A, B], b monadic.Data[any, aType[A]]) monadic.Data[any, aType[A]] {
	if IsLeft(a) {
		return Left[any](*a.left)
	}
	eb := b.(Either[A, any])
	if IsLeft(eb) {
		return Left[any](*eb.left)
	}
	return Right[A](f(*a.right, *eb.right))
}

func (Either[A, B]) Bind(ma Either[A, B], mm func(B) monadic.Data[any, aType[A]]) monadic.Data[any, aType[A]] {
	if IsRight(ma) {
		return mm(*ma.right)
	}
	return Left[any](*ma.left)
}

func (Either[A, B]) Do(proc func() Either[A, B]) (ret Either[A, B]) {
	defer func() {
		e := recover()
		if a, ok := e.(aType[A]); ok {
			ret = Left[B](a.e)
			return
		}
		panic(e)
	}()
	ret = proc()
	return
}

func (m Either[A, B]) X() B {
	if IsRight(m) {
		return *m.right
	}
	panic(aType[A]{*m.left})
}
