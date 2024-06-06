package lazy

import (
	"sync"

	"github.com/arcane-craft/monadic"
)

type Value[A any] func() A

func New[A any](f func() A) Value[A] {
	var once sync.Once
	var ret A
	return func() A {
		once.Do(func() {
			ret = f()
		})
		return ret
	}
}

func Zero[A any]() Value[A] {
	return func() A {
		var zero A
		return zero
	}
}

func Pure[A any](v A) Value[A] {
	return func() A {
		return v
	}
}

type Delayable[A any] interface {
	Delay() Value[A]
}

func ImplDelayable[T Delayable[T]]() monadic.Unit {
	return monadic.Unit{}
}

func Delay[A any](v A) Value[A] {
	panic(monadic.NotSupportForTest)
}
