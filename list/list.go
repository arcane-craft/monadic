package list

import (
	"github.com/arcane-craft/monadic/algebra"
	"github.com/arcane-craft/monadic/foldable"
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

var _ = monad.ImplMonad[List[any]]()
var _ = algebra.ImplMonoid[List[any]]()
var _ = foldable.ImplFoldable[List[any]]()
