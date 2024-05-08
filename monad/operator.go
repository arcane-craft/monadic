package monad

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
)

type MonadClass[M monadic.Data[A, _E], A any, _E monadic.Nillable] interface {
	applicative.ApplicativeClass[M, A, _E]
}

func ImplMonadClass[
	M interface {
		monadic.Data[A, _E]
		MonadClass[M, A, _E]
	},
	A any,
	_E monadic.Nillable]() monadic.Void {
	return monadic.Void{}
}

type MonadDoClass[M monadic.Data[A, _E], A any, _E monadic.Nillable] interface {
	MonadClass[M, A, _E]
	X() A
}

func ImplMonadDoClass[
	M interface {
		monadic.Data[A, _E]
		MonadDoClass[M, A, _E]
	},
	A any,
	_E monadic.Nillable]() monadic.Void {
	return monadic.Void{}
}

func Return[
	M interface {
		monadic.Data[A, _E]
		MonadClass[M, A, _E]
	},
	A any,
	_E monadic.Nillable](a A) M {
	return applicative.Pure[M](a)
}

func Bind[
	MA interface {
		monadic.Data[A, _E]
		MonadClass[MA, A, _E]
	},
	MB interface {
		monadic.Data[B, _E]
		MonadClass[MB, B, _E]
	},
	A, B any,
	_E monadic.Nillable](ma MA, mm func(A) MB) MB {
	var mb MB
	return mb.Init(func() MB {
		r, e := ma.Resolve()
		if !e.IsNil() {
			return mb.Throw(e)
		}
		return mm(r)
	})
}

func Then[
	MA interface {
		monadic.Data[A, _E]
		MonadClass[MA, A, _E]
	},
	MB interface {
		monadic.Data[B, _E]
		MonadClass[MB, B, _E]
	},
	A, B any,
	_E monadic.Nillable](ma MA, mb func() MB) MB {
	return Bind(ma, func(A) MB {
		return mb()
	})
}

func Join[
	MM interface {
		monadic.Data[M, _E]
		MonadClass[MM, M, _E]
	},
	M interface {
		monadic.Data[A, _E]
		MonadClass[M, A, _E]
	},
	A any,
	_E monadic.Nillable](mm MM) M {
	return Bind(mm, func(m M) M {
		return m
	})
}

type doException[E monadic.Nillable] struct {
	e E
}

func Do[
	M interface {
		monadic.Data[A, _E]
		MonadDoClass[M, A, _E]
	},
	A any,
	_E monadic.Nillable](proc func() M) M {
	var m M
	return m.Init(func() (out M) {
		defer func() {
			e := recover()
			if e != nil {
				if doE, ok := e.(*doException[_E]); ok {
					out = m.Throw(doE.e)
					return
				}
			}
			panic(e)
		}()
		out = proc()
		return
	})
}

func X[
	M interface {
		monadic.Data[A, _E]
		MonadDoClass[M, A, _E]
	},
	A any,
	_E monadic.Nillable](m M) A {
	r, e := m.Resolve()
	if !e.IsNil() {
		panic(&doException[_E]{
			e: e,
		})
	}
	return r
}

func LiftM[
	MB interface {
		monadic.Data[B, _E]
		MonadClass[MB, B, _E]
	},
	A, B any,
	MA interface {
		monadic.Data[A, _E]
		MonadClass[MA, A, _E]
	},
	_E monadic.Nillable](f func(A) B, ma MA) MB {
	return Bind(ma, func(a A) MB {
		var mb MB
		return mb.Init(func() MB {
			return mb.Pure(f(a))
		})
	})
}

func LiftM2[
	MC interface {
		monadic.Data[C, _E]
		MonadClass[MC, C, _E]
	},
	A, B, C any,
	MA interface {
		monadic.Data[A, _E]
		MonadClass[MA, A, _E]
	},
	MB interface {
		monadic.Data[B, _E]
		MonadClass[MB, B, _E]
	},
	_E monadic.Nillable](f func(A, B) C, ma MA, mb MB) MC {
	return Bind(ma, func(a A) MC {
		return Bind(mb, func(b B) MC {
			var mc MC
			return mc.Init(func() MC {
				return mc.Pure(f(a, b))
			})
		})
	})
}

func LiftM3[
	MD interface {
		monadic.Data[D, _E]
		MonadClass[MD, D, _E]
	},
	A, B, C, D any,
	MA interface {
		monadic.Data[A, _E]
		MonadClass[MA, A, _E]
	},
	MB interface {
		monadic.Data[B, _E]
		MonadClass[MB, B, _E]
	},
	MC interface {
		monadic.Data[C, _E]
		MonadClass[MC, C, _E]
	},
	_E monadic.Nillable](f func(A, B, C) D, ma MA, mb MB, mc MC) MD {
	return Bind(ma, func(a A) MD {
		return Bind(mb, func(b B) MD {
			return Bind(mc, func(c C) MD {
				var md MD
				return md.Init(func() MD {
					return md.Pure(f(a, b, c))
				})
			})
		})
	})
}

func LiftM4[
	ME interface {
		monadic.Data[E, _E]
		MonadClass[ME, E, _E]
	},
	A, B, C, D, E any,
	MA interface {
		monadic.Data[A, _E]
		MonadClass[MA, A, _E]
	},
	MB interface {
		monadic.Data[B, _E]
		MonadClass[MB, B, _E]
	},
	MC interface {
		monadic.Data[C, _E]
		MonadClass[MC, C, _E]
	},
	MD interface {
		monadic.Data[D, _E]
		MonadClass[MD, D, _E]
	},
	_E monadic.Nillable](f func(A, B, C, D) E, ma MA, mb MB, mc MC, md MD) ME {
	return Bind(ma, func(a A) ME {
		return Bind(mb, func(b B) ME {
			return Bind(mc, func(c C) ME {
				return Bind(md, func(d D) ME {
					var me ME
					return me.Init(func() ME {
						return me.Pure(f(a, b, c, d))
					})
				})
			})
		})
	})
}

func LiftM5[
	MF interface {
		monadic.Data[F, _E]
		MonadClass[MF, F, _E]
	},
	A, B, C, D, E, F any,
	MA interface {
		monadic.Data[A, _E]
		MonadClass[MA, A, _E]
	},
	MB interface {
		monadic.Data[B, _E]
		MonadClass[MB, B, _E]
	},
	MC interface {
		monadic.Data[C, _E]
		MonadClass[MC, C, _E]
	},
	MD interface {
		monadic.Data[D, _E]
		MonadClass[MD, D, _E]
	},
	ME interface {
		monadic.Data[E, _E]
		MonadClass[ME, E, _E]
	},
	_E monadic.Nillable](f func(A, B, C, D, E) F, ma MA, mb MB, mc MC, md MD, me ME) MF {
	return Bind(ma, func(a A) MF {
		return Bind(mb, func(b B) MF {
			return Bind(mc, func(c C) MF {
				return Bind(md, func(d D) MF {
					return Bind(me, func(e E) MF {
						var mf MF
						return mf.Init(func() MF {
							return mf.Pure(f(a, b, c, d, e))
						})
					})
				})
			})
		})
	})
}
