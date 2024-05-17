package option

import (
	"github.com/arcane-craft/monadic"
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

func (Option[A]) Traverse() {}
