package option

import "github.com/arcane-craft/monadic/lazy"

type Option[A any] struct {
	v     lazy.Value[A]
	valid lazy.Bool
}

func Some[A any](a lazy.Value[A]) lazy.Value[Option[A]] {
	return lazy.Const(Option[A]{
		v:     a,
		valid: lazy.Const(true),
	})
}

func None[A any]() lazy.Value[Option[A]] {
	return lazy.Const(Option[A]{
		valid: lazy.Const(false),
	})
}

func IsNone[A any](v lazy.Value[Option[A]]) lazy.Bool {
	return lazy.Bind(v, func(o Option[A]) lazy.Bool {
		return lazy.Not(o.valid)
	})
}

func Map[A, B any](fa lazy.Value[Option[A]], m func(lazy.Value[A]) lazy.Value[B]) lazy.Value[Option[B]] {
	return lazy.Bind(fa, func(o Option[A]) lazy.Value[Option[B]] {
		return lazy.IfThenElse(o.valid,
			func() lazy.Value[Option[B]] {
				return Some(m(o.v))
			},
			func() lazy.Value[Option[B]] {
				return None[B]()
			},
		)
	})
}

func Bind[A, B any](ma lazy.Value[Option[A]], mm func(lazy.Value[A]) lazy.Value[Option[B]]) lazy.Value[Option[B]] {
	return lazy.Bind(ma, func(o Option[A]) lazy.Value[Option[B]] {
		return lazy.IfThenElse(o.valid,
			func() lazy.Value[Option[B]] {
				return mm(o.v)
			},
			func() lazy.Value[Option[B]] {
				return None[B]()
			},
		)
	})
}
