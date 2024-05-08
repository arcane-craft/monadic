package option

import (
	"github.com/arcane-craft/monadic/monad"
)

func (m Option[A]) Resolve() (A, eType) {
	var zero A
	if m.v == nil {
		return zero, eType(true)
	}
	return *m.v, eType(false)
}

func (Option[A]) Throw(eType) Option[A] {
	return None[A]()
}

func (Option[A]) Init(f func() Option[A]) Option[A] {
	return f()
}

func (Option[A]) Pure(a A) Option[A] {
	return Some(a)
}

func (m Option[A]) X() A {
	return monad.X(m)
}
