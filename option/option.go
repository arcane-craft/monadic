package option

import "github.com/arcane-craft/monadic/lazy"

type Option[A any] struct {
	v     lazy.Value[A]
	valid bool
}

func Some[A any](a lazy.Value[A]) lazy.Value[Option[A]] {
	return lazy.Const(Option[A]{
		v:     a,
		valid: true,
	})
}

func None[A any]() lazy.Value[Option[A]] {
	return lazy.Const(Option[A]{
		valid: false,
	})
}

func IsNone[A any](v lazy.Value[Option[A]]) lazy.Bool {
	return lazy.Map(v, func(o Option[A]) bool {
		return !o.valid
	})
}

func Map[A, B any](fa lazy.Value[Option[A]], m func(lazy.Value[A]) lazy.Value[B]) lazy.Value[Option[B]] {
	return lazy.Bind(fa, func(o Option[A]) lazy.Value[Option[B]] {
		if o.valid {
			return Some(m(o.v))
		}
		return None[B]()
	})
}

func Bind[A, B any](ma lazy.Value[Option[A]], mm func(lazy.Value[A]) lazy.Value[Option[B]]) lazy.Value[Option[B]] {
	return lazy.Bind(ma, func(o Option[A]) lazy.Value[Option[B]] {
		if o.valid {
			return mm(o.v)
		}
		return None[B]()
	})
}
