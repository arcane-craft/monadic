package option

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/foldable"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/lazy"
)

func (Option[A]) Kind() aType {
	return aType{}
}

func (Option[A]) Concretize(o monadic.Data[any, aType]) Option[A] {
	oi := o.(Option[any])
	if oi.v != nil {
		return Some((*oi.v).(A))
	}
	return None[A]()
}

func (Option[A]) Abstract(o Option[A]) monadic.Data[any, aType] {
	if o.v != nil {
		return Some(any(*o.v))
	}
	return None[any]()
}

func (o Option[A]) Delay() lazy.Value[Option[A]] {
	panic(monadic.NotSupportForTest)
}

func (Option[A]) Map(m func(A) any, fa Option[A]) monadic.Data[any, aType] {
	if IsSome(fa) {
		return Some(m(*fa.v))
	}
	return None[any]()
}

func (Option[A]) Pure(a A) Option[A] {
	return Some(a)
}

func (Option[A]) LiftA2(f func(A, any) any, a Option[A], b monadic.Data[any, aType]) monadic.Data[any, aType] {
	if IsNone(a) {
		return None[any]()
	}
	eb := b.(Option[any])
	if IsNone(eb) {
		return None[any]()
	}
	return Some(f(*a.v, *eb.v))
}

func (Option[A]) Empty() Option[A] {
	return None[A]()
}

func (Option[A]) Or(a Option[A], b Option[A]) Option[A] {
	if a.v != nil {
		return Some(*a.v)
	}
	return b
}

func (Option[A]) Bind(ma Option[A], mm func(A) monadic.Data[any, aType]) monadic.Data[any, aType] {
	if ma.v != nil {
		return mm(*ma.v)
	}
	return None[any]()
}

func (Option[A]) Do(proc func() Option[A]) (ret Option[A]) {
	defer func() {
		e := recover()
		if _, ok := e.(aType); ok {
			ret = None[A]()
			return
		}
		panic(e)
	}()
	ret = proc()
	return
}

func (m Option[A]) X() A {
	if m.v != nil {
		return *m.v
	}
	panic(aType{})
}

func (Option[A]) Append(a Option[A], b Option[A]) Option[A] {
	if IsNone(a) {
		return b
	}
	return a
}

func (Option[A]) Neutral() Option[A] {
	return None[A]()
}

func (Option[A]) Foldr(f func(A, any) any, init any, input Option[A]) any {
	if IsNone(input) {
		return init
	}
	return f(*input.v, init)
}

func Traverse[
	FTB interface {
		applicative.Applicative[FTB, Option[B], _F]
		monadic.Data[Option[B], _F]
	},
	A any,
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[B, _F]
	},
	B any,
	_F any,
](f func(A) FB, t Option[A]) FTB {
	if IsNone(t) {
		return applicative.Pure[FTB](None[B]())
	}
	return applicative.LiftA[FTB](func(b B) Option[B] {
		return Some(b)
	}, f(FromSome(t)))
}

func Sequence[
	FTA interface {
		applicative.Applicative[FTA, Option[A], _F]
		monadic.Data[Option[A], _F]
	},
	FA interface {
		applicative.Applicative[FA, A, _F]
		monadic.Data[A, _F]
	},
	A any,
	_F any,
](t Option[FA]) FTA {
	return Traverse[FTA](basics.Id, t)
}

func For[
	FTB interface {
		applicative.Applicative[FTB, Option[B], _F]
		monadic.Data[Option[B], _F]
	},
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[B, _F]
	},
	A, B any,
	_F any,
](t Option[A], f func(A) FB) FTB {
	return Traverse[FTB](f, t)
}

func Traverse_[
	FN interface {
		applicative.Applicative[FN, monadic.Unit, _F]
		monadic.Data[monadic.Unit, _F]
	},
	A any,
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[FB, _F]
	},
	B any,
	_F any,
](f func(A) FB, t Option[A]) FN {
	return foldable.Foldr(function.Uncurry(function.Compose(function.Curry(applicative.ApplyR[FB, FN]), f)), applicative.Pure[FN](monadic.Unit{}), t)
}

func Sequence_[
	FN interface {
		applicative.Applicative[FN, monadic.Unit, _F]
		monadic.Data[monadic.Unit, _F]
	},
	FA interface {
		applicative.Applicative[FA, A, _F]
		monadic.Data[A, _F]
	},
	A any,
	_F any,
](t Option[FA]) FN {
	return foldable.Foldr(applicative.ApplyR, applicative.Pure[FN](monadic.Unit{}), t)
}

func For_[
	FN interface {
		applicative.Applicative[FN, monadic.Unit, _F]
		monadic.Data[monadic.Unit, _F]
	},
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[FB, _F]
	},
	A, B any,
	_F any,
](t Option[A], f func(A) FB) FN {
	return Traverse_[FN](f, t)
}
