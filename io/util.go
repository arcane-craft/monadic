package io

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/result"
	"github.com/arcane-craft/monadic/tuple"
)

func ApplySucc(f func()) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			f()
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func LiftSucc(f func()) func() IO[monadic.Void] {
	return func() IO[monadic.Void] {
		return ApplySucc(f)
	}
}

func ApplySuccX[C any](ctx *Context[C], f func()) {
	Continue(ctx, LiftSucc(f)())
}

func LiftSuccX[C any](f func()) func(*Context[C]) {
	return func(ctx *Context[C]) {
		ApplySuccX(ctx, f)
	}
}

func Apply1PSucc[P any](f func(P), p lazy.Value[P]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			f(lazy.Eval(p))
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift1PSucc[P any](f func(P)) func(lazy.Value[P]) IO[monadic.Void] {
	return func(p lazy.Value[P]) IO[monadic.Void] {
		return Apply1PSucc(f, p)
	}
}

func Apply1PSuccX[C, P any](ctx *Context[C], f func(P), p lazy.Value[P]) {
	Continue(ctx, Lift1PSucc(f)(p))
}

func Lift1PSuccX[C, P any](f func(P)) func(*Context[C], lazy.Value[P]) {
	return func(ctx *Context[C], p lazy.Value[P]) {
		Apply1PSuccX(ctx, f, p)
	}
}

func Apply2PSucc[P1, P2 any](f func(P1, P2), p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			f(lazy.Eval(p1), lazy.Eval(p2))
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift2PSucc[P1, P2 any](f func(P1, P2)) func(lazy.Value[P1], lazy.Value[P2]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[monadic.Void] {
		return Apply2PSucc(f, p1, p2)
	}
}

func Apply2PSuccX[C, P1, P2 any](ctx *Context[C], f func(P1, P2), p1 lazy.Value[P1], p2 lazy.Value[P2]) {
	Continue(ctx, Lift2PSucc(f)(p1, p2))
}

func Lift2PSuccX[C, P1, P2 any](f func(P1, P2)) func(*Context[C], lazy.Value[P1], lazy.Value[P2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2]) {
		Apply2PSuccX(ctx, f, p1, p2)
	}
}

func Apply3PSucc[P1, P2, P3 any](f func(P1, P2, P3), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3))
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift3PSucc[P1, P2, P3 any](f func(P1, P2, P3)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[monadic.Void] {
		return Apply3PSucc(f, p1, p2, p3)
	}
}

func Apply3PSuccX[C, P1, P2, P3 any](ctx *Context[C], f func(P1, P2, P3), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) {
	Continue(ctx, Lift3PSucc(f)(p1, p2, p3))
}

func Lift3PSuccX[C, P1, P2, P3 any](f func(P1, P2, P3)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) {
		Apply3PSuccX(ctx, f, p1, p2, p3)
	}
}

func Apply4PSucc[P1, P2, P3, P4 any](f func(P1, P2, P3, P4), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4))
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift4PSucc[P1, P2, P3, P4 any](f func(P1, P2, P3, P4)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[monadic.Void] {
		return Apply4PSucc(f, p1, p2, p3, p4)
	}
}

func Apply4PSuccX[C, P1, P2, P3, P4 any](ctx *Context[C], f func(P1, P2, P3, P4), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) {
	Continue(ctx, Lift4PSucc(f)(p1, p2, p3, p4))
}

func Lift4PSuccX[C, P1, P2, P3, P4 any](f func(P1, P2, P3, P4)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) {
		Apply4PSuccX(ctx, f, p1, p2, p3, p4)
	}
}

func Apply5PSucc[P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4), lazy.Eval(p5))
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift5PSucc[P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[monadic.Void] {
		return Apply5PSucc(f, p1, p2, p3, p4, p5)
	}
}

func Apply5PSuccX[C, P1, P2, P3, P4, P5 any](ctx *Context[C], f func(P1, P2, P3, P4, P5), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) {
	Continue(ctx, Lift5PSucc(f)(p1, p2, p3, p4, p5))
}

func Lift5PSuccX[C, P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) {
		Apply5PSuccX(ctx, f, p1, p2, p3, p4, p5)
	}
}

func ApplyVarPSucc[P any](f func(...P), v ...lazy.Value[P]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			var ps []P
			for _, p := range v {
				ps = append(ps, lazy.Eval(p))
			}
			f(ps...)
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func LiftVarPSucc[P any](f func(...P)) func(...lazy.Value[P]) IO[monadic.Void] {
	return func(a ...lazy.Value[P]) IO[monadic.Void] {
		return ApplyVarPSucc(f, a...)
	}
}

func ApplyVarPSuccX[C, P any](ctx *Context[C], f func(...P), p ...lazy.Value[P]) {
	Continue(ctx, LiftVarPSucc(f)(p...))
}

func LiftVarPSuccX[C, P any](f func(...P)) func(*Context[C], ...lazy.Value[P]) {
	return func(ctx *Context[C], p ...lazy.Value[P]) {
		ApplyVarPSuccX(ctx, f, p...)
	}
}

func Apply1RSucc[R any](f func() R) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r := f()
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift1RSucc[R any](f func() R) func() IO[R] {
	return func() IO[R] {
		return Apply1RSucc(f)
	}
}

func Apply1RSuccX[C, R any](ctx *Context[C], f func() R) lazy.Value[R] {
	return From(ctx, Lift1RSucc(f)())
}

func Lift1RSuccX[C, R any](f func() R) func(*Context[C]) lazy.Value[R] {
	return func(ctx *Context[C]) lazy.Value[R] {
		return Apply1RSuccX(ctx, f)
	}
}

func Apply1P1RSucc[P, R any](f func(P) R, p lazy.Value[P]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r := f(lazy.Eval(p))
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift1P1RSucc[P, R any](f func(P) R) func(lazy.Value[P]) IO[R] {
	return func(p lazy.Value[P]) IO[R] {
		return Apply1P1RSucc(f, p)
	}
}

func Apply1P1RSuccX[C, P, R any](ctx *Context[C], f func(P) R, p lazy.Value[P]) lazy.Value[R] {
	return From(ctx, Lift1P1RSucc(f)(p))
}

func Lift1P1RSuccX[C, P, R any](f func(P) R) func(*Context[C], lazy.Value[P]) lazy.Value[R] {
	return func(ctx *Context[C], p lazy.Value[P]) lazy.Value[R] {
		return Apply1P1RSuccX(ctx, f, p)
	}
}

func Apply2P1RSucc[P1, P2, R any](f func(P1, P2) R, p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r := f(lazy.Eval(p1), lazy.Eval(p2))
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift2P1RSucc[P1, P2, R any](f func(P1, P2) R) func(lazy.Value[P1], lazy.Value[P2]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[R] {
		return Apply2P1RSucc(f, p1, p2)
	}
}

func Apply2P1RSuccX[C, P1, P2, R any](ctx *Context[C], f func(P1, P2) R, p1 lazy.Value[P1], p2 lazy.Value[P2]) lazy.Value[R] {
	return From(ctx, Lift2P1RSucc(f)(p1, p2))
}

func Lift2P1RSuccX[C, P1, P2, R any](f func(P1, P2) R) func(*Context[C], lazy.Value[P1], lazy.Value[P2]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2]) lazy.Value[R] {
		return Apply2P1RSuccX(ctx, f, p1, p2)
	}
}

func Apply3P1RSucc[P1, P2, P3, R any](f func(P1, P2, P3) R, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3))
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift3P1RSucc[P1, P2, P3, R any](f func(P1, P2, P3) R) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[R] {
		return Apply3P1RSucc(f, p1, p2, p3)
	}
}

func Apply3P1RSuccX[C, P1, P2, P3, R any](ctx *Context[C], f func(P1, P2, P3) R, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) lazy.Value[R] {
	return From(ctx, Lift3P1RSucc(f)(p1, p2, p3))
}

func Lift3P1RSuccX[C, P1, P2, P3, R any](f func(P1, P2, P3) R) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) lazy.Value[R] {
		return Apply3P1RSuccX(ctx, f, p1, p2, p3)
	}
}

func Apply4P1RSucc[P1, P2, P3, P4, R any](f func(P1, P2, P3, P4) R, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4))
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift4P1RSucc[P1, P2, P3, P4, R any](f func(P1, P2, P3, P4) R) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[R] {
		return Apply4P1RSucc(f, p1, p2, p3, p4)
	}
}

func Apply4P1RSuccX[C, P1, P2, P3, P4, R any](ctx *Context[C], f func(P1, P2, P3, P4) R, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) lazy.Value[R] {
	return From(ctx, Lift4P1RSucc(f)(p1, p2, p3, p4))
}

func Lift4P1RSuccX[C, P1, P2, P3, P4, R any](f func(P1, P2, P3, P4) R) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) lazy.Value[R] {
		return Apply4P1RSuccX(ctx, f, p1, p2, p3, p4)
	}
}

func Apply5P1RSucc[P1, P2, P3, P4, P5, R any](f func(P1, P2, P3, P4, P5) R, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4), lazy.Eval(p5))
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift5P1RSucc[P1, P2, P3, P4, P5, R any](f func(P1, P2, P3, P4, P5) R) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[R] {
		return Apply5P1RSucc(f, p1, p2, p3, p4, p5)
	}
}

func Apply5P1RSuccX[C, P1, P2, P3, P4, P5, R any](ctx *Context[C], f func(P1, P2, P3, P4, P5) R, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) lazy.Value[R] {
	return From(ctx, Lift5P1RSucc(f)(p1, p2, p3, p4, p5))
}

func Lift5P1RSuccX[C, P1, P2, P3, P4, P5, R any](f func(P1, P2, P3, P4, P5) R) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) lazy.Value[R] {
		return Apply5P1RSuccX(ctx, f, p1, p2, p3, p4, p5)
	}
}

func Apply2RSucc[R1, R2 any](f func() (R1, R2)) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2 := f()
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift2RSucc[R1, R2 any](f func() (R1, R2)) func() IO[tuple.Tuple2[R1, R2]] {
	return func() IO[tuple.Tuple2[R1, R2]] {
		return Apply2RSucc(f)
	}
}

func Apply2RSuccX[C, R1, R2 any](ctx *Context[C], f func() (R1, R2)) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift2RSucc(f)()), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift2RSuccX[C, R1, R2 any](f func() (R1, R2)) func(*Context[C]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply2RSuccX(ctx, f)
	}
}

func Apply1P2RSucc[P, R1, R2 any](f func(P) (R1, R2), p lazy.Value[P]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2 := f(lazy.Eval(p))
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift1P2RSucc[P, R1, R2 any](f func(P) (R1, R2)) func(lazy.Value[P]) IO[tuple.Tuple2[R1, R2]] {
	return func(p lazy.Value[P]) IO[tuple.Tuple2[R1, R2]] {
		return Apply1P2RSucc(f, p)
	}
}

func Apply1P2RSuccX[C, P, R1, R2 any](ctx *Context[C], f func(P) (R1, R2), p lazy.Value[P]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift1P2RSucc(f)(p)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift1P2RSuccX[C, P, R1, R2 any](f func(P) (R1, R2)) func(*Context[C], lazy.Value[P]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p lazy.Value[P]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply1P2RSuccX(ctx, f, p)
	}
}

func Apply2P2RSucc[P1, P2, R1, R2 any](f func(P1, P2) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2 := f(lazy.Eval(p1), lazy.Eval(p2))
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift2P2RSucc[P1, P2, R1, R2 any](f func(P1, P2) (R1, R2)) func(lazy.Value[P1], lazy.Value[P2]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[tuple.Tuple2[R1, R2]] {
		return Apply2P2RSucc(f, p1, p2)
	}
}

func Apply2P2RSuccX[C, P1, P2, R1, R2 any](ctx *Context[C], f func(P1, P2) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift2P2RSucc(f)(p1, p2)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift2P2RSuccX[C, P1, P2, R1, R2 any](f func(P1, P2) (R1, R2)) func(*Context[C], lazy.Value[P1], lazy.Value[P2]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply2P2RSuccX(ctx, f, p1, p2)
	}
}

func Apply3P2RSucc[P1, P2, P3, R1, R2 any](f func(P1, P2, P3) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2 := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3))
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift3P2RSucc[P1, P2, P3, R1, R2 any](f func(P1, P2, P3) (R1, R2)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[tuple.Tuple2[R1, R2]] {
		return Apply3P2RSucc(f, p1, p2, p3)
	}
}

func Apply3P2RSuccX[C, P1, P2, P3, R1, R2 any](ctx *Context[C], f func(P1, P2, P3) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift3P2RSucc(f)(p1, p2, p3)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift3P2RSuccX[C, P1, P2, P3, R1, R2 any](f func(P1, P2, P3) (R1, R2)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply3P2RSuccX(ctx, f, p1, p2, p3)
	}
}

func Apply4P2RSucc[P1, P2, P3, P4, R1, R2 any](f func(P1, P2, P3, P4) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2 := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4))
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift4P2RSucc[P1, P2, P3, P4, R1, R2 any](f func(P1, P2, P3, P4) (R1, R2)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[tuple.Tuple2[R1, R2]] {
		return Apply4P2RSucc(f, p1, p2, p3, p4)
	}
}

func Apply4P2RSuccX[C, P1, P2, P3, P4, R1, R2 any](ctx *Context[C], f func(P1, P2, P3, P4) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift4P2RSucc(f)(p1, p2, p3, p4)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift4P2RSuccX[C, P1, P2, P3, P4, R1, R2 any](f func(P1, P2, P3, P4) (R1, R2)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply4P2RSuccX(ctx, f, p1, p2, p3, p4)
	}
}

func Apply5P2RSucc[P1, P2, P3, P4, P5, R1, R2 any](f func(P1, P2, P3, P4, P5) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2 := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4), lazy.Eval(p5))
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift5P2RSucc[P1, P2, P3, P4, P5, R1, R2 any](f func(P1, P2, P3, P4, P5) (R1, R2)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[tuple.Tuple2[R1, R2]] {
		return Apply5P2RSucc(f, p1, p2, p3, p4, p5)
	}
}

func Apply5P2RSuccX[C, P1, P2, P3, P4, P5, R1, R2 any](ctx *Context[C], f func(P1, P2, P3, P4, P5) (R1, R2), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift5P2RSucc(f)(p1, p2, p3, p4, p5)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift5P2RSuccX[C, P1, P2, P3, P4, P5, R1, R2 any](f func(P1, P2, P3, P4, P5) (R1, R2)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply5P2RSuccX(ctx, f, p1, p2, p3, p4, p5)
	}
}

func Apply(f func() error) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			err := f()
			if err != nil {
				return lazy.Eval(result.Fail[monadic.Void](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift(f func() error) func() IO[monadic.Void] {
	return func() IO[monadic.Void] {
		return Apply(f)
	}
}

func ApplyX[C any](ctx *Context[C], f func() error) {
	Continue(ctx, Lift(f)())
}

func LiftX[C any](f func() error) func(*Context[C]) {
	return func(ctx *Context[C]) {
		ApplyX(ctx, f)
	}
}

func Apply1P[P any](f func(P) error, p lazy.Value[P]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			err := f(lazy.Eval(p))
			if err != nil {
				return lazy.Eval(result.Fail[monadic.Void](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift1P[P any](f func(P) error) func(lazy.Value[P]) IO[monadic.Void] {
	return func(p lazy.Value[P]) IO[monadic.Void] {
		return Apply1P(f, p)
	}
}

func Apply1PX[C, P any](ctx *Context[C], f func(P) error, p lazy.Value[P]) {
	Continue(ctx, Lift1P(f)(p))
}

func Lift1PX[C, P any](f func(P) error) func(*Context[C], lazy.Value[P]) {
	return func(ctx *Context[C], p lazy.Value[P]) {
		Apply1PX(ctx, f, p)
	}
}

func Apply2P[P1, P2 any](f func(P1, P2) error, p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			err := f(lazy.Eval(p1), lazy.Eval(p2))
			if err != nil {
				return lazy.Eval(result.Fail[monadic.Void](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift2P[P1, P2 any](f func(P1, P2) error) func(lazy.Value[P1], lazy.Value[P2]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[monadic.Void] {
		return Apply2P(f, p1, p2)
	}
}

func Apply2PX[C, P1, P2 any](ctx *Context[C], f func(P1, P2) error, p1 lazy.Value[P1], p2 lazy.Value[P2]) {
	Continue(ctx, Lift2P(f)(p1, p2))
}

func Lift2PX[C, P1, P2 any](f func(P1, P2) error) func(*Context[C], lazy.Value[P1], lazy.Value[P2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2]) {
		Apply2PX(ctx, f, p1, p2)
	}
}

func Apply3P[P1, P2, P3 any](f func(P1, P2, P3) error, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3))
			if err != nil {
				return lazy.Eval(result.Fail[monadic.Void](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift3P[P1, P2, P3 any](f func(P1, P2, P3) error) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[monadic.Void] {
		return Apply3P(f, p1, p2, p3)
	}
}

func Apply3PX[C, P1, P2, P3 any](ctx *Context[C], f func(P1, P2, P3) error, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) {
	Continue(ctx, Lift3P(f)(p1, p2, p3))
}

func Lift3PX[C, P1, P2, P3 any](f func(P1, P2, P3) error) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) {
		Apply3PX(ctx, f, p1, p2, p3)
	}
}

func Apply4P[P1, P2, P3, P4 any](f func(P1, P2, P3, P4) error, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4))
			if err != nil {
				return lazy.Eval(result.Fail[monadic.Void](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift4P[P1, P2, P3, P4 any](f func(P1, P2, P3, P4) error) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[monadic.Void] {
		return Apply4P(f, p1, p2, p3, p4)
	}
}

func Apply4PX[C, P1, P2, P3, P4 any](ctx *Context[C], f func(P1, P2, P3, P4) error, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) {
	Continue(ctx, Lift4P(f)(p1, p2, p3, p4))
}

func Lift4PX[C, P1, P2, P3, P4 any](f func(P1, P2, P3, P4) error) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) {
		Apply4PX(ctx, f, p1, p2, p3, p4)
	}
}

func Apply5P[P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5) error, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4), lazy.Eval(p5))
			if err != nil {
				return lazy.Eval(result.Fail[monadic.Void](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func Lift5P[P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5) error) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) IO[monadic.Void] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[monadic.Void] {
		return Apply5P(f, p1, p2, p3, p4, p5)
	}
}

func Apply5PX[C, P1, P2, P3, P4, P5 any](ctx *Context[C], f func(P1, P2, P3, P4, P5) error, p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) {
	Continue(ctx, Lift5P(f)(p1, p2, p3, p4, p5))
}

func Lift5PX[C, P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5) error) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) {
		Apply5PX(ctx, f, p1, p2, p3, p4, p5)
	}
}

func ApplyVarP[P any](f func(...P) error, v ...lazy.Value[P]) IO[monadic.Void] {
	return IO[monadic.Void]{
		v: lazy.New(func() either.Either[monadic.Void, error] {
			var ps []P
			for _, p := range v {
				ps = append(ps, lazy.Eval(p))
			}
			err := f(ps...)
			if err != nil {
				return lazy.Eval(result.Fail[monadic.Void](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(monadic.Void{})))
		}),
	}
}

func LiftVarP[P any](f func(...P) error) func(...lazy.Value[P]) IO[monadic.Void] {
	return func(a ...lazy.Value[P]) IO[monadic.Void] {
		return ApplyVarP(f, a...)
	}
}

func ApplyVarPX[C, P any](ctx *Context[C], f func(...P) error, p ...lazy.Value[P]) {
	Continue(ctx, LiftVarP(f)(p...))
}

func LiftVarPX[C, P any](f func(...P) error) func(*Context[C], ...lazy.Value[P]) {
	return func(ctx *Context[C], p ...lazy.Value[P]) {
		ApplyVarPX(ctx, f, p...)
	}
}

func Apply1R[R any](f func() (R, error)) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r, err := f()
			if err != nil {
				return lazy.Eval(result.Fail[R](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift1R[R any](f func() (R, error)) func() IO[R] {
	return func() IO[R] {
		return Apply1R(f)
	}
}

func Apply1RX[C, R any](ctx *Context[C], f func() (R, error)) lazy.Value[R] {
	return From(ctx, Lift1R(f)())
}

func Lift1RX[C, R any](f func() (R, error)) func(*Context[C]) lazy.Value[R] {
	return func(ctx *Context[C]) lazy.Value[R] {
		return Apply1RX(ctx, f)
	}
}

func Apply1P1R[P, R any](f func(P) (R, error), p lazy.Value[P]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r, err := f(lazy.Eval(p))
			if err != nil {
				return lazy.Eval(result.Fail[R](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift1P1R[P, R any](f func(P) (R, error)) func(lazy.Value[P]) IO[R] {
	return func(p lazy.Value[P]) IO[R] {
		return Apply1P1R(f, p)
	}
}

func Apply1P1RX[C, P, R any](ctx *Context[C], f func(P) (R, error), p lazy.Value[P]) lazy.Value[R] {
	return From(ctx, Lift1P1R(f)(p))
}

func Lift1P1RX[C, P, R any](f func(P) (R, error)) func(*Context[C], lazy.Value[P]) lazy.Value[R] {
	return func(ctx *Context[C], p lazy.Value[P]) lazy.Value[R] {
		return Apply1P1RX(ctx, f, p)
	}
}

func Apply2P1R[P1, P2, R any](f func(P1, P2) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r, err := f(lazy.Eval(p1), lazy.Eval(p2))
			if err != nil {
				return lazy.Eval(result.Fail[R](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift2P1R[P1, P2, R any](f func(P1, P2) (R, error)) func(lazy.Value[P1], lazy.Value[P2]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[R] {
		return Apply2P1R(f, p1, p2)
	}
}

func Apply2P1RX[C, P1, P2, R any](ctx *Context[C], f func(P1, P2) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2]) lazy.Value[R] {
	return From(ctx, Lift2P1R(f)(p1, p2))
}

func Lift2P1RX[C, P1, P2, R any](f func(P1, P2) (R, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2]) lazy.Value[R] {
		return Apply2P1RX(ctx, f, p1, p2)
	}
}

func Apply3P1R[P1, P2, P3, R any](f func(P1, P2, P3) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r, err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3))
			if err != nil {
				return lazy.Eval(result.Fail[R](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift3P1R[P1, P2, P3, R any](f func(P1, P2, P3) (R, error)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[R] {
		return Apply3P1R(f, p1, p2, p3)
	}
}

func Apply3P1RX[C, P1, P2, P3, R any](ctx *Context[C], f func(P1, P2, P3) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) lazy.Value[R] {
	return From(ctx, Lift3P1R(f)(p1, p2, p3))
}

func Lift3P1RX[C, P1, P2, P3, R any](f func(P1, P2, P3) (R, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) lazy.Value[R] {
		return Apply3P1RX(ctx, f, p1, p2, p3)
	}
}

func Apply4P1R[P1, P2, P3, P4, R any](f func(P1, P2, P3, P4) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r, err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4))
			if err != nil {
				return lazy.Eval(result.Fail[R](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift4P1R[P1, P2, P3, P4, R any](f func(P1, P2, P3, P4) (R, error)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[R] {
		return Apply4P1R(f, p1, p2, p3, p4)
	}
}

func Apply4P1RX[C, P1, P2, P3, P4, R any](ctx *Context[C], f func(P1, P2, P3, P4) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) lazy.Value[R] {
	return From(ctx, Lift4P1R(f)(p1, p2, p3, p4))
}

func Lift4P1RX[C, P1, P2, P3, P4, R any](f func(P1, P2, P3, P4) (R, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) lazy.Value[R] {
		return Apply4P1RX(ctx, f, p1, p2, p3, p4)
	}
}

func Apply5P1R[P1, P2, P3, P4, P5, R any](f func(P1, P2, P3, P4, P5) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			r, err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4), lazy.Eval(p5))
			if err != nil {
				return lazy.Eval(result.Fail[R](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func Lift5P1R[P1, P2, P3, P4, P5, R any](f func(P1, P2, P3, P4, P5) (R, error)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) IO[R] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[R] {
		return Apply5P1R(f, p1, p2, p3, p4, p5)
	}
}

func Apply5P1RX[C, P1, P2, P3, P4, P5, R any](ctx *Context[C], f func(P1, P2, P3, P4, P5) (R, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) lazy.Value[R] {
	return From(ctx, Lift5P1R(f)(p1, p2, p3, p4, p5))
}

func Lift5P1RX[C, P1, P2, P3, P4, P5, R any](f func(P1, P2, P3, P4, P5) (R, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) lazy.Value[R] {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) lazy.Value[R] {
		return Apply5P1RX(ctx, f, p1, p2, p3, p4, p5)
	}
}

func ApplyVarP1R[P, R any](f func(...P) (R, error), v ...lazy.Value[P]) IO[R] {
	return IO[R]{
		v: lazy.New(func() either.Either[R, error] {
			var ps []P
			for _, p := range v {
				ps = append(ps, lazy.Eval(p))
			}
			r, err := f(ps...)
			if err != nil {
				return lazy.Eval(result.Fail[R](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(r)))
		}),
	}
}

func LiftVarP1R[P, R any](f func(...P) (R, error)) func(...lazy.Value[P]) IO[R] {
	return func(a ...lazy.Value[P]) IO[R] {
		return ApplyVarP1R(f, a...)
	}
}

func ApplyVarP1RX[C, P, R any](ctx *Context[C], f func(...P) (R, error), p ...lazy.Value[P]) lazy.Value[R] {
	return From(ctx, LiftVarP1R(f)(p...))
}

func LiftVarP1RX[C, P, R any](f func(...P) (R, error)) func(*Context[C], ...lazy.Value[P]) lazy.Value[R] {
	return func(ctx *Context[C], p ...lazy.Value[P]) lazy.Value[R] {
		return ApplyVarP1RX(ctx, f, p...)
	}
}

func Apply2R[R1, R2 any](f func() (R1, R2, error)) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2, err := f()
			if err != nil {
				return lazy.Eval(result.Fail[tuple.Tuple2[R1, R2]](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift2R[R1, R2 any](f func() (R1, R2, error)) func() IO[tuple.Tuple2[R1, R2]] {
	return func() IO[tuple.Tuple2[R1, R2]] {
		return Apply2R(f)
	}
}

func Apply2RX[C, R1, R2 any](ctx *Context[C], f func() (R1, R2, error)) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift2R(f)()), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift2RX[C, R1, R2 any](f func() (R1, R2, error)) func(*Context[C]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply2RX(ctx, f)
	}
}

func Apply1P2R[P, R1, R2 any](f func(P) (R1, R2, error), p lazy.Value[P]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2, err := f(lazy.Eval(p))
			if err != nil {
				return lazy.Eval(result.Fail[tuple.Tuple2[R1, R2]](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift1P2R[P, R1, R2 any](f func(P) (R1, R2, error)) func(lazy.Value[P]) IO[tuple.Tuple2[R1, R2]] {
	return func(p lazy.Value[P]) IO[tuple.Tuple2[R1, R2]] {
		return Apply1P2R(f, p)
	}
}

func Apply1P2RX[C, P, R1, R2 any](ctx *Context[C], f func(P) (R1, R2, error), p lazy.Value[P]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift1P2R(f)(p)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift1P2RX[C, P, R1, R2 any](f func(P) (R1, R2, error)) func(*Context[C], lazy.Value[P]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p lazy.Value[P]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply1P2RX(ctx, f, p)
	}
}

func Apply2P2R[P1, P2, R1, R2 any](f func(P1, P2) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2, err := f(lazy.Eval(p1), lazy.Eval(p2))
			if err != nil {
				return lazy.Eval(result.Fail[tuple.Tuple2[R1, R2]](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift2P2R[P1, P2, R1, R2 any](f func(P1, P2) (R1, R2, error)) func(lazy.Value[P1], lazy.Value[P2]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2]) IO[tuple.Tuple2[R1, R2]] {
		return Apply2P2R(f, p1, p2)
	}
}

func Apply2P2RX[C, P1, P2, R1, R2 any](ctx *Context[C], f func(P1, P2) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift2P2R(f)(p1, p2)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift2P2RX[C, P1, P2, R1, R2 any](f func(P1, P2) (R1, R2, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply2P2RX(ctx, f, p1, p2)
	}
}

func Apply3P2R[P1, P2, P3, R1, R2 any](f func(P1, P2, P3) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2, err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3))
			if err != nil {
				return lazy.Eval(result.Fail[tuple.Tuple2[R1, R2]](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift3P2R[P1, P2, P3, R1, R2 any](f func(P1, P2, P3) (R1, R2, error)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) IO[tuple.Tuple2[R1, R2]] {
		return Apply3P2R(f, p1, p2, p3)
	}
}

func Apply3P2RX[C, P1, P2, P3, R1, R2 any](ctx *Context[C], f func(P1, P2, P3) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift3P2R(f)(p1, p2, p3)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift3P2RX[C, P1, P2, P3, R1, R2 any](f func(P1, P2, P3) (R1, R2, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply3P2RX(ctx, f, p1, p2, p3)
	}
}

func Apply4P2R[P1, P2, P3, P4, R1, R2 any](f func(P1, P2, P3, P4) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2, err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4))
			if err != nil {
				return lazy.Eval(result.Fail[tuple.Tuple2[R1, R2]](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift4P2R[P1, P2, P3, P4, R1, R2 any](f func(P1, P2, P3, P4) (R1, R2, error)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) IO[tuple.Tuple2[R1, R2]] {
		return Apply4P2R(f, p1, p2, p3, p4)
	}
}

func Apply4P2RX[C, P1, P2, P3, P4, R1, R2 any](ctx *Context[C], f func(P1, P2, P3, P4) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift4P2R(f)(p1, p2, p3, p4)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift4P2RX[C, P1, P2, P3, P4, R1, R2 any](f func(P1, P2, P3, P4) (R1, R2, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply4P2RX(ctx, f, p1, p2, p3, p4)
	}
}

func Apply5P2R[P1, P2, P3, P4, P5, R1, R2 any](f func(P1, P2, P3, P4, P5) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[tuple.Tuple2[R1, R2]] {
	return IO[tuple.Tuple2[R1, R2]]{
		v: lazy.New(func() either.Either[tuple.Tuple2[R1, R2], error] {
			r1, r2, err := f(lazy.Eval(p1), lazy.Eval(p2), lazy.Eval(p3), lazy.Eval(p4), lazy.Eval(p5))
			if err != nil {
				return lazy.Eval(result.Fail[tuple.Tuple2[R1, R2]](lazy.Const(err)))
			}
			return lazy.Eval(result.Ok(lazy.Const(tuple.New2(r1, r2))))
		}),
	}
}

func Lift5P2R[P1, P2, P3, P4, P5, R1, R2 any](f func(P1, P2, P3, P4, P5) (R1, R2, error)) func(lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) IO[tuple.Tuple2[R1, R2]] {
		return Apply5P2R(f, p1, p2, p3, p4, p5)
	}
}

func Apply5P2RX[C, P1, P2, P3, P4, P5, R1, R2 any](ctx *Context[C], f func(P1, P2, P3, P4, P5) (R1, R2, error), p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) (lazy.Value[R1], lazy.Value[R2]) {
	return lazy.Eval(lazy.Map(From(ctx, Lift5P2R(f)(p1, p2, p3, p4, p5)), func(t tuple.Tuple2[R1, R2]) tuple.Tuple2[lazy.Value[R1], lazy.Value[R2]] {
		return tuple.Map2(t, func(r1 R1, r2 R2) (lazy.Value[R1], lazy.Value[R2]) {
			return lazy.New(func() R1 { return r1 }), lazy.New(func() R2 { return r2 })
		})
	}))()
}

func Lift5P2RX[C, P1, P2, P3, P4, P5, R1, R2 any](f func(P1, P2, P3, P4, P5) (R1, R2, error)) func(*Context[C], lazy.Value[P1], lazy.Value[P2], lazy.Value[P3], lazy.Value[P4], lazy.Value[P5]) (lazy.Value[R1], lazy.Value[R2]) {
	return func(ctx *Context[C], p1 lazy.Value[P1], p2 lazy.Value[P2], p3 lazy.Value[P3], p4 lazy.Value[P4], p5 lazy.Value[P5]) (lazy.Value[R1], lazy.Value[R2]) {
		return Apply5P2RX(ctx, f, p1, p2, p3, p4, p5)
	}
}
