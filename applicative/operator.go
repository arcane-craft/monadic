package applicative

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/functor"
)

type ApplicativeClass[F monadic.Data[A, _E], A any, _E monadic.Nillable] interface {
	functor.FunctorClass[F, A, _E]
}

func ImplApplicativeClass[
	F interface {
		monadic.Data[A, _E]
		ApplicativeClass[F, A, _E]
	},
	A any,
	_E monadic.Nillable]() monadic.Void {
	return monadic.Void{}
}

type AlternativeClass[F monadic.Data[A, _E], A any, _E monadic.Nillable] interface {
	ApplicativeClass[F, A, _E]
	Empty() F
	Or(F) F
}

func ImplAlternativeClass[
	F interface {
		monadic.Data[A, _E]
		AlternativeClass[F, A, _E]
	},
	A any,
	_E monadic.Nillable]() monadic.Void {
	return monadic.Void{}
}

func Pure[
	F interface {
		monadic.Data[A, _E]
		ApplicativeClass[F, A, _E]
	},
	A any,
	_E monadic.Nillable](a A) F {
	var f F
	return f.Init(func() F {
		return f.Pure(a)
	})
}

func SeqApply[
	FB interface {
		monadic.Data[B, _E]
		ApplicativeClass[FB, B, _E]
	},
	A, B any,
	FM interface {
		monadic.Data[func(A) B, _E]
		ApplicativeClass[FM, func(A) B, _E]
	},
	FA interface {
		monadic.Data[A, _E]
		ApplicativeClass[FA, A, _E]
	},
	_E monadic.Nillable](fm FM, fa FA) FB {
	var fb FB
	return fb.Init(func() FB {
		m, e := fm.Resolve()
		if !e.IsNil() {
			return fb.Throw(e)
		}
		a, e := fa.Resolve()
		if !e.IsNil() {
			return fb.Throw(e)
		}
		return fb.Pure(m(a))
	})
}

func LiftA2[
	FC interface {
		monadic.Data[C, _E]
		ApplicativeClass[FC, C, _E]
	},
	A, B, C any,
	FA interface {
		monadic.Data[A, _E]
		ApplicativeClass[FA, A, _E]
	},
	FB interface {
		monadic.Data[B, _E]
		ApplicativeClass[FB, B, _E]
	},
	_E monadic.Nillable](f func(A, B) C, fa FA, fb FB) FC {
	var fc FC
	return fc.Init(func() FC {
		a, e := fa.Resolve()
		if !e.IsNil() {
			return fc.Throw(e)
		}
		b, e := fb.Resolve()
		if !e.IsNil() {
			return fc.Throw(e)
		}
		return fc.Pure(f(a, b))
	})
}

func Empty[
	F interface {
		monadic.Data[A, _E]
		AlternativeClass[F, A, _E]
	},
	A any,
	_E monadic.Nillable]() F {
	var f F
	return f.Empty()
}

func Or[
	F interface {
		monadic.Data[A, _E]
		AlternativeClass[F, A, _E]
	},
	A any,
	_E monadic.Nillable](a F, b F) F {
	return a.Or(b)
}
