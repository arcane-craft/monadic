package functor

import "github.com/arcane-craft/monadic"

type FunctorClass[F monadic.Data[A, _E], A any, _E monadic.Nillable] interface {
	Throw(_E) F
	Init(func() F) F
	Pure(A) F
}

func ImplFunctorClass[
	F interface {
		monadic.Data[A, _E]
		FunctorClass[F, A, _E]
	},
	A any,
	_E monadic.Nillable]() monadic.Void {
	return monadic.Void{}
}

func Map[
	FB interface {
		monadic.Data[B, _E]
		FunctorClass[FB, B, _E]
	},
	A, B any,
	FA interface {
		monadic.Data[A, _E]
		FunctorClass[FA, A, _E]
	},
	_E monadic.Nillable](m func(A) B, fa FA) FB {
	var fb FB
	return fb.Init(func() FB {
		r, e := fa.Resolve()
		if !e.IsNil() {
			return fb.Throw(e)
		}
		return fb.Pure(m(r))
	})
}

func Replace[
	FA interface {
		monadic.Data[A, _E]
		FunctorClass[FA, A, _E]
	},
	FB interface {
		monadic.Data[B, _E]
		FunctorClass[FB, B, _E]
	},
	A, B any,
	_E monadic.Nillable](a A, fb FB) FA {
	return Map[FA](func(B) A {
		return a
	}, fb)
}
