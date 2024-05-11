package list

import "github.com/arcane-craft/monadic"

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

func (List[A]) Apply(fm monadic.Data[func(A) any, aType], fa List[A]) monadic.Data[any, aType] {
	fmi := fm.(rList[func(A) any, aType])
	ret := make(List[any], 0, len(fmi)+len(fa))
	for _, m := range fmi {
		for _, a := range fa {
			ret = append(ret, m(a))
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
	panic("not support for testing, please optimize first")
}

func (m List[A]) X() A {
	panic("not support for testing, please optimize first")
}
