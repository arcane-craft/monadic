package monad

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/functor"
)

type MonadClass[
	M monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	Bind(M, func(A) monadic.Data[any, _T]) monadic.Data[any, _T]
}

func ImplMonadClass[
	M interface {
		MonadClass[M, A, _T]
		applicative.ApplicativeClass[M, A, _T]
		functor.FunctorClass[M, A, _T]
		monadic.Generalize[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Void {
	return monadic.Void{}
}

type MonadDoClass[
	M monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	MonadClass[M, A, _T]
	Do(func() M) M
	X() A
}

func ImplMonadDoClass[
	M interface {
		MonadDoClass[M, A, _T]
		applicative.ApplicativeClass[M, A, _T]
		functor.FunctorClass[M, A, _T]
		monadic.Generalize[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Void {
	return monadic.Void{}
}

func Bind[
	MA interface {
		MonadClass[MA, A, _T]
		applicative.ApplicativeClass[MA, A, _T]
		functor.FunctorClass[MA, A, _T]
		monadic.Generalize[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		MonadClass[MB, B, _T]
		applicative.ApplicativeClass[MB, B, _T]
		functor.FunctorClass[MB, B, _T]
		monadic.Generalize[MB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](ma MA, mm func(A) MB) MB {
	return monadic.Zero[MB]().Concretize(monadic.Zero[MA]().Bind(ma, func(a A) monadic.Data[any, _T] {
		return monadic.Zero[MB]().Abstract(mm(a))
	}))
}

func Then[
	MA interface {
		MonadClass[MA, A, _T]
		applicative.ApplicativeClass[MA, A, _T]
		functor.FunctorClass[MA, A, _T]
		monadic.Generalize[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		MonadClass[MB, B, _T]
		applicative.ApplicativeClass[MB, B, _T]
		functor.FunctorClass[MB, B, _T]
		monadic.Generalize[MB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](ma MA, mb func() MB) MB {
	return Bind(ma, func(A) MB {
		return mb()
	})
}

func Join[
	MM interface {
		MonadClass[MM, M, _T]
		applicative.ApplicativeClass[MM, M, _T]
		functor.FunctorClass[MM, M, _T]
		monadic.Generalize[MM, M, _T]
		monadic.Data[A, _T]
	},
	M interface {
		MonadClass[M, A, _T]
		applicative.ApplicativeClass[M, A, _T]
		functor.FunctorClass[M, A, _T]
		monadic.Generalize[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](mm MM) M {
	return Bind(mm, monadic.Id)
}

func DoInit[
	M interface {
		MonadDoClass[M, A, _T]
		applicative.ApplicativeClass[M, A, _T]
		functor.FunctorClass[M, A, _T]
		monadic.Generalize[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() M {
	return Return[M](monadic.Zero[A]())
}

func Do[
	M interface {
		MonadDoClass[M, A, _T]
		applicative.ApplicativeClass[M, A, _T]
		functor.FunctorClass[M, A, _T]
		monadic.Generalize[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](proc func() M) M {
	return monadic.Zero[M]().Do(proc)
}

func LiftM[
	MB interface {
		MonadClass[MB, B, _T]
		applicative.ApplicativeClass[MB, B, _T]
		functor.FunctorClass[MB, B, _T]
		monadic.Generalize[MB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	MA interface {
		MonadClass[MA, A, _T]
		applicative.ApplicativeClass[MA, A, _T]
		functor.FunctorClass[MA, A, _T]
		monadic.Generalize[MA, A, _T]
		monadic.Data[A, _T]
	},
	_T any,
](f func(A) B, ma MA) MB {
	return functor.Map[MB](f, ma)
}

func LiftM2[
	MC interface {
		MonadClass[MC, C, _T]
		applicative.ApplicativeClass[MC, C, _T]
		functor.FunctorClass[MC, C, _T]
		monadic.Generalize[MC, C, _T]
		monadic.Data[C, _T]
	},
	A, B, C any,
	MA interface {
		MonadClass[MA, A, _T]
		applicative.ApplicativeClass[MA, A, _T]
		functor.FunctorClass[MA, A, _T]
		monadic.Generalize[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		MonadClass[MB, B, _T]
		applicative.ApplicativeClass[MB, B, _T]
		functor.FunctorClass[MB, B, _T]
		monadic.Generalize[MB, B, _T]
		monadic.Data[B, _T]
	},
	_T any,
](f func(A, B) C, ma MA, mb MB) MC {
	return Bind(ma, func(a A) MC {
		return LiftM[MC](function.Partial(f, a), mb)
	})
}

func LiftM3[
	MD interface {
		MonadClass[MD, D, _T]
		applicative.ApplicativeClass[MD, D, _T]
		functor.FunctorClass[MD, D, _T]
		monadic.Generalize[MD, D, _T]
		monadic.Data[D, _T]
	},
	A, B, C, D any,
	MA interface {
		MonadClass[MA, A, _T]
		applicative.ApplicativeClass[MA, A, _T]
		functor.FunctorClass[MA, A, _T]
		monadic.Generalize[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		MonadClass[MB, B, _T]
		applicative.ApplicativeClass[MB, B, _T]
		functor.FunctorClass[MB, B, _T]
		monadic.Generalize[MB, B, _T]
		monadic.Data[B, _T]
	},
	MC interface {
		MonadClass[MC, C, _T]
		applicative.ApplicativeClass[MC, C, _T]
		functor.FunctorClass[MC, C, _T]
		monadic.Generalize[MC, C, _T]
		monadic.Data[C, _T]
	},
	_T any,
](f func(A, B, C) D, ma MA, mb MB, mc MC) MD {
	return Bind(ma, func(a A) MD {
		return LiftM2[MD](function.Partial3(f, a), mb, mc)
	})
}

func LiftM4[
	ME interface {
		MonadClass[ME, E, _T]
		applicative.ApplicativeClass[ME, E, _T]
		functor.FunctorClass[ME, E, _T]
		monadic.Generalize[ME, E, _T]
		monadic.Data[E, _T]
	},
	A, B, C, D, E any,
	MA interface {
		MonadClass[MA, A, _T]
		applicative.ApplicativeClass[MA, A, _T]
		functor.FunctorClass[MA, A, _T]
		monadic.Generalize[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		MonadClass[MB, B, _T]
		applicative.ApplicativeClass[MB, B, _T]
		functor.FunctorClass[MB, B, _T]
		monadic.Generalize[MB, B, _T]
		monadic.Data[B, _T]
	},
	MC interface {
		MonadClass[MC, C, _T]
		applicative.ApplicativeClass[MC, C, _T]
		functor.FunctorClass[MC, C, _T]
		monadic.Generalize[MC, C, _T]
		monadic.Data[C, _T]
	},
	MD interface {
		MonadClass[MD, D, _T]
		applicative.ApplicativeClass[MD, D, _T]
		functor.FunctorClass[MD, D, _T]
		monadic.Generalize[MD, D, _T]
		monadic.Data[D, _T]
	},
	_T any,
](f func(A, B, C, D) E, ma MA, mb MB, mc MC, md MD) ME {
	return Bind(ma, func(a A) ME {
		return LiftM3[ME](function.Partial4(f, a), mb, mc, md)
	})
}

func LiftM5[
	MF interface {
		MonadClass[MF, F, _T]
		applicative.ApplicativeClass[MF, F, _T]
		functor.FunctorClass[MF, F, _T]
		monadic.Generalize[MF, F, _T]
		monadic.Data[F, _T]
	},
	A, B, C, D, E, F any,
	MA interface {
		MonadClass[MA, A, _T]
		applicative.ApplicativeClass[MA, A, _T]
		functor.FunctorClass[MA, A, _T]
		monadic.Generalize[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		MonadClass[MB, B, _T]
		applicative.ApplicativeClass[MB, B, _T]
		functor.FunctorClass[MB, B, _T]
		monadic.Generalize[MB, B, _T]
		monadic.Data[B, _T]
	},
	MC interface {
		MonadClass[MC, C, _T]
		applicative.ApplicativeClass[MC, C, _T]
		functor.FunctorClass[MC, C, _T]
		monadic.Generalize[MC, C, _T]
		monadic.Data[C, _T]
	},
	MD interface {
		MonadClass[MD, D, _T]
		applicative.ApplicativeClass[MD, D, _T]
		functor.FunctorClass[MD, D, _T]
		monadic.Generalize[MD, D, _T]
		monadic.Data[D, _T]
	},
	ME interface {
		MonadClass[ME, E, _T]
		applicative.ApplicativeClass[ME, E, _T]
		functor.FunctorClass[ME, E, _T]
		monadic.Generalize[ME, E, _T]
		monadic.Data[E, _T]
	},
	_T any,
](f func(A, B, C, D, E) F, ma MA, mb MB, mc MC, md MD, me ME) MF {
	return Bind(ma, func(a A) MF {
		return LiftM4[MF](function.Partial5(f, a), mb, mc, md, me)
	})
}

func Return[
	M interface {
		MonadClass[M, A, _T]
		applicative.ApplicativeClass[M, A, _T]
		functor.FunctorClass[M, A, _T]
		monadic.Generalize[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](a A) M {
	return monadic.Zero[M]().Pure(a)
}
