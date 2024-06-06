package monad

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/foldable"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/functor"
)

type Monad[
	M monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	applicative.Applicative[M, A, _T]
	Bind(M, func(A) monadic.Data[any, _T]) monadic.Data[any, _T]
}

func ImplMonad[
	M interface {
		Monad[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Unit {
	return monadic.Unit{}
}

type MonadDo[
	M monadic.Data[A, _T],
	A any,
	_T any,
] interface {
	Monad[M, A, _T]
	Do(func() M) M
	X() A
}

func ImplMonadDo[
	M interface {
		MonadDo[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() monadic.Unit {
	return monadic.Unit{}
}

func Bind[
	MA interface {
		Monad[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		Monad[MB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](ma MA, mm func(A) MB) MB {
	return basics.Zero[MB]().Concretize(basics.Zero[MA]().Bind(ma, func(a A) monadic.Data[any, _T] {
		return basics.Zero[MB]().Abstract(mm(a))
	}))
}

func Then[
	MA interface {
		Monad[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		Monad[MB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](ma MA, mb func() MB) MB {
	return Bind(ma, func(A) MB {
		return mb()
	})
}

func Then_[
	MA interface {
		Monad[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		Monad[MB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	_T any,
](ma MA, mb MB) MB {
	return Then(ma, func() MB {
		return mb
	})
}

func Join[
	MM interface {
		Monad[MM, M, _T]
		monadic.Data[A, _T]
	},
	M interface {
		Monad[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](mm MM) M {
	return Bind(mm, basics.Id)
}

func DoInit[
	M interface {
		MonadDo[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
]() M {
	return applicative.Pure[M](basics.Zero[A]())
}

func Do[
	M interface {
		MonadDo[M, A, _T]
		monadic.Data[A, _T]
	},
	A any,
	_T any,
](proc func() M) M {
	return basics.Zero[M]().Do(proc)
}

func DoF[
	M interface {
		MonadDo[M, A, _T]
		monadic.Data[A, _T]
	},
	P, A any,
	_T any,
](proc func(P) M) func(P) M {
	return func(p P) M {
		return basics.Zero[M]().Do(func() M {
			return proc(p)
		})
	}
}

func DoF2[
	M interface {
		MonadDo[M, A, _T]
		monadic.Data[A, _T]
	},
	P1, P2, A any,
	_T any,
](proc func(P1, P2) M) func(P1, P2) M {
	return func(p1 P1, p2 P2) M {
		return basics.Zero[M]().Do(func() M {
			return proc(p1, p2)
		})
	}
}

func DoF3[
	M interface {
		MonadDo[M, A, _T]
		monadic.Data[A, _T]
	},
	P1, P2, P3, A any,
	_T any,
](proc func(P1, P2, P3) M) func(P1, P2, P3) M {
	return func(p1 P1, p2 P2, p3 P3) M {
		return basics.Zero[M]().Do(func() M {
			return proc(p1, p2, p3)
		})
	}
}

func LiftM[
	MB interface {
		Monad[MB, B, _T]
		monadic.Data[B, _T]
	},
	A, B any,
	MA interface {
		Monad[MA, A, _T]
		monadic.Data[A, _T]
	},
	_T any,
](f func(A) B, ma MA) MB {
	return functor.Map[MB](f, ma)
}

func LiftM2[
	MC interface {
		Monad[MC, C, _T]
		monadic.Data[C, _T]
	},
	A, B, C any,
	MA interface {
		Monad[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		Monad[MB, B, _T]
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
		Monad[MD, D, _T]
		monadic.Data[D, _T]
	},
	A, B, C, D any,
	MA interface {
		Monad[MA, A, _T]
		monadic.Data[A, _T]
	},
	MB interface {
		Monad[MB, B, _T]
		monadic.Data[B, _T]
	},
	MC interface {
		Monad[MC, C, _T]
		monadic.Data[C, _T]
	},
	_T any,
](f func(A, B, C) D, ma MA, mb MB, mc MC) MD {
	return Bind(ma, func(a A) MD {
		return LiftM2[MD](function.Partial3(f, a), mb, mc)
	})
}

func FoldlM[
	TB interface {
		foldable.Foldable[TB, B, _T]
		monadic.Data[B, _T]
	},
	MA interface {
		Monad[MA, A, _M]
		monadic.Data[A, _M]
	},
	A, B any,
	_M, _T any,
](fm func(A, B) MA, a0 A, x TB) MA {
	return foldable.Foldl(func(ma MA, b B) MA {
		return Bind(ma, func(a A) MA {
			return function.Flip(fm)(b, a)
		})
	}, applicative.Pure[MA](a0), x)
}
