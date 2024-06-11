package list

import (
	"github.com/arcane-craft/monadic/algebra"
	"github.com/arcane-craft/monadic/foldable"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/monad"
)

type aType struct{}

type rList[A any, _T any] []A

type List[A any] rList[A, aType]

func Cons[A any](a A, l List[A]) List[A] {
	return append(l, a)
}

func Nil[A any]() List[A] {
	return nil
}

func L[A any](a ...A) List[A] {
	return append(make(List[A], 0, len(a)), a...)
}

func IsCons[A any](l List[A]) bool {
	return len(l) > 0
}

func IsNil[A any](l List[A]) bool {
	return len(l) <= 0
}

func Head[A any](l List[A]) A {
	if len(l) <= 0 {
		panic("empty list")
	}
	return l[0]
}

func Tail[A any](l List[A]) List[A] {
	if len(l) <= 0 {
		panic("empty list")
	}
	return l[1:]
}

func Take[A any](n int, l List[A]) List[A] {
	if n >= len(l) {
		return l
	}
	return l[:n]
}

func Drop[A any](n int, l List[A]) List[A] {
	if n >= len(l) {
		return Nil[A]()
	}
	return l[n:]
}

func ChunksOf[A any](n int, l List[A]) List[List[A]] {
	ret := Nil[List[A]]()
	for len(l) > 0 {
		ret = Cons(Take(n, l), ret)
		l = Drop(n, l)
	}
	return ret
}

func Filter[A any](p func(A) bool, l List[A]) List[A] {
	var ret List[A]
	for _, e := range l {
		if p(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

var _ = lazy.ImplDelayable[List[any]]()
var _ = monad.ImplMonad[List[any]]()
var _ = algebra.ImplMonoid[List[any]]()
var _ = foldable.ImplFoldable[List[any]]()
