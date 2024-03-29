package io

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/result"
	"github.com/arcane-craft/monadic/tuple"
)

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
