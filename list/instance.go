package list

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/applicative"
	"github.com/arcane-craft/monadic/basics"
	"github.com/arcane-craft/monadic/foldable"
	"github.com/arcane-craft/monadic/function"
	"github.com/arcane-craft/monadic/lazy"
)

func (List[A]) Kind() aType {
	return aType{}
}

func (List[A]) Concretize(o monadic.Data[any, aType]) List[A] {
	oi := o.(List[any])
	ret := make(List[A], 0, len(oi))
	for _, a := range oi {
		ret = append(ret, a.(A))
	}
	return ret
}

func (List[A]) Abstract(o List[A]) monadic.Data[any, aType] {
	ret := make(List[any], 0, len(o))
	for _, a := range o {
		ret = append(ret, any(a))
	}
	return ret
}

func (l List[A]) Delay() lazy.Value[List[A]] {
	panic(monadic.NotSupportForTest)
}

func (List[A]) Map(m func(A) any, fa List[A]) monadic.Data[any, aType] {
	ret := make(List[any], 0, len(fa))
	for _, a := range fa {
		ret = append(ret, m(a))
	}
	return ret
}

func (List[A]) Pure(a A) List[A] {
	return List[A]{a}
}

func (List[A]) LiftA2(f func(A, any) any, a List[A], b monadic.Data[any, aType]) monadic.Data[any, aType] {
	bb := b.(List[any])
	ret := make(List[any], 0, len(a)+len(bb))
	for _, ea := range a {
		for _, eb := range bb {
			ret = append(ret, f(ea, eb))
		}
	}
	return ret
}

func (List[A]) Empty() List[A] {
	return List[A]{}
}

func (List[A]) Or(a List[A], b List[A]) List[A] {
	ret := make(List[A], 0, len(a)+len(b))
	ret = append(ret, a...)
	ret = append(ret, b...)
	return ret
}

func (List[A]) Bind(ma List[A], mm func(A) monadic.Data[any, aType]) monadic.Data[any, aType] {
	ret := make(List[any], 0, len(ma))
	for _, a := range ma {
		ret = append(ret, mm(a).(List[any])...)
	}
	return ret
}

func (List[A]) Append(a List[A], b List[A]) List[A] {
	return append(append(make(List[A], 0, len(a)+len(b)), a...), b...)
}

func (List[A]) Neutral() List[A] {
	return List[A]{}
}

func (List[A]) Foldr(f func(A, any) any, init any, input List[A]) any {
	ret := init
	for i := len(input) - 1; i >= 0; i-- {
		ret = f(input[i], ret)
	}
	return ret
}

func Traverse[
	FTB interface {
		applicative.Applicative[FTB, List[B], _F]
		monadic.Data[List[B], _F]
	},
	A any,
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[B, _F]
	},
	B any,
	_F any,
](f func(A) FB, t List[A]) FTB {
	if IsNil(t) {
		return applicative.Pure[FTB](Nil[B]())
	}
	return applicative.LiftA2[FTB](func(b B, l List[B]) List[B] {
		return Cons(b, l)
	}, f(Head(t)), Traverse[FTB](f, Tail(t)))
}

func Sequence[
	FTA interface {
		applicative.Applicative[FTA, List[A], _F]
		monadic.Data[List[A], _F]
	},
	FA interface {
		applicative.Applicative[FA, A, _F]
		monadic.Data[A, _F]
	},
	A any,
	_F any,
](t List[FA]) FTA {
	return Traverse[FTA](basics.Id, t)
}

func For[
	FTB interface {
		applicative.Applicative[FTB, List[B], _F]
		monadic.Data[List[B], _F]
	},
	FB interface {
		applicative.Applicative[FB, B, _F]
		monadic.Data[B, _F]
	},
	A, B any,
	_F any,
](t List[A], f func(A) FB) FTB {
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
](f func(A) FB, t List[A]) FN {
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
](t List[FA]) FN {
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
](t List[A], f func(A) FB) FN {
	return Traverse_[FN](f, t)
}
