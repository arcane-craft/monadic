package io

import (
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/result"
)

type IO[A any] struct {
	v lazy.Value[either.Either[A, error]]
}

func Ret[A any](v lazy.Value[A]) IO[A] {
	return IO[A]{
		v: lazy.Const(result.Ok(v)),
	}
}

func Bind[A, B any](ma IO[A], mm func(lazy.Value[A]) IO[B]) IO[B] {
	return IO[B]{
		v: result.Bind(ma.v, func(v lazy.Value[A]) lazy.Value[either.Either[B, error]] {
			return mm(v).v
		}),
	}
}

func Then[A, B any](ma IO[A], mb func() IO[B]) IO[B] {
	return IO[B]{
		v: result.Bind(ma.v, func(v lazy.Value[A]) lazy.Value[either.Either[B, error]] {
			return mb().v
		}),
	}
}

type Context[A any] func(func() IO[A]) IO[A]

func From[A, B any](ctx *Context[B], m IO[A]) (ret lazy.Value[A]) {
	var cacheRet lazy.Value[A]
	ret = lazy.New(func() A {
		return lazy.Eval(cacheRet)
	})
	prevCtx := *ctx
	*ctx = func(step func() IO[B]) IO[B] {
		return prevCtx(func() IO[B] {
			return Bind(m, func(v lazy.Value[A]) IO[B] {
				cacheRet = v
				return step()
			})
		})
	}
	return
}

func Continue[A, B any](ctx *Context[B], m IO[A]) {
	prevCtx := *ctx
	*ctx = func(step func() IO[B]) IO[B] {
		return prevCtx(func() IO[B] {
			return Then(m, step)
		})
	}
}

func Do[A any](block func(*Context[A]) IO[A]) IO[A] {
	ctx := Context[A](func(step func() IO[A]) IO[A] {
		return step()
	})
	ret := block(&ctx)
	return ctx(func() IO[A] {
		return ret
	})
}

func Try[A any](m IO[A]) IO[either.Either[A, error]] {
	return IO[either.Either[A, error]]{
		v: lazy.Const(either.Left[either.Either[A, error], error](m.v)),
	}
}

func Descript[A any](m IO[A], desc lazy.Value[string]) IO[A] {
	return IO[A]{
		v: lazy.Bind(result.IsFail(m.v), func(fail bool) lazy.Value[either.Either[A, error]] {
			if fail {
				return result.WrapFail(lazy.Map(desc, func(s string) string {
					return s + " failed:"
				}), m.v)
			}
			return m.v
		}),
	}
}
