package list

import (
	"github.com/arcane-craft/monadic"
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

func (List[A]) Do(proc func() List[A]) List[A] {
	panic(monadic.NotSupportForTest)
}

func (m List[A]) X() A {
	panic(monadic.NotSupportForTest)
}

func (List[A]) Append(a List[A], b List[A]) List[A] {
	return append(append(make(List[A], 0, len(a)+len(b)), a...), b...)
}

func (List[A]) Neutral() List[A] {
	return List[A]{}
}

func (List[A]) Foldr(f func(A, any) any, init any, input List[A]) any {
	ret := init
	for _, x := range input {
		ret = f(x, ret)
	}
	return ret
}
