package option

import "github.com/arcane-craft/monadic/lazy"

type Option[A any] struct {
	v     lazy.Value[A]
	valid lazy.Value[bool]
}

func Some[A any](a lazy.Value[A]) Option[A] {
	return Option[A]{
		v:     a,
		valid: lazy.Const(true),
	}
}

func None[A any]() Option[A] {
	return Option[A]{
		valid: lazy.Const(false),
	}
}

func IsNone[A any](v lazy.Value[Option[A]]) lazy.Value[bool] {
	return lazy.Map(v, func(o Option[A]) bool {
		return !lazy.Eval(o.valid)
	})
}

func Map[A, B any](fa lazy.Value[Option[A]], m func(lazy.Value[A]) lazy.Value[B]) lazy.Value[Option[B]] {
	return lazy.Map(fa, func(o Option[A]) Option[B] {
		if !lazy.Eval(o.valid) {
			return None[B]()
		}
		return Some(m(o.v))
	})
}

func Bind[A, B any](ma lazy.Value[Option[A]], mm func(lazy.Value[A]) lazy.Value[Option[B]]) lazy.Value[Option[B]] {
	return lazy.Bind(ma, func(o Option[A]) lazy.Value[Option[B]] {
		if !lazy.Eval(o.valid) {
			return lazy.Const(None[B]())
		}
		return mm(o.v)
	})
}
