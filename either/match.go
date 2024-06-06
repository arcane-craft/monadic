package either

import "github.com/arcane-craft/monadic/function"

func Match[R, A, B any](left func(A) R, right func(B) R, e Either[A, B]) R {
	if e.left != nil {
		return left(*e.left)
	}
	return right(*e.right)
}

func MatchF[R, A, B any](left func(A) R, right func(B) R) func(Either[A, B]) R {
	return function.Partial32(Match, left, right)
}

func MatchLeft[A, B any](e Either[A, B]) A {
	if !IsLeft(e) {
		panic("MatchLeft on right value of Either")
	}
	return *e.left
}

func MatchRight[A, B any](e Either[A, B]) A {
	if !IsRight(e) {
		panic("MatchLeft() on Right value of Either")
	}
	return *e.left
}
