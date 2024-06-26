package option

import "github.com/arcane-craft/monadic/function"

func Match[R, A any](some func(A) R, none func() R, o Option[A]) R {
	if IsSome(o) {
		return some(*o.v)
	}
	return none()
}

func MatchF[R, A any](some func(A) R, none func() R) func(Option[A]) R {
	return function.Partial32(Match, some, none)
}

func MatchSome[A any](o Option[A]) A {
	if !IsSome(o) {
		panic("MatchSome() on None value of Option")
	}
	return *o.v
}
