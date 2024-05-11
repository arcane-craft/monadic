package list

import (
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

var _ = monad.ImplMonadDoClass[List[any]]()
