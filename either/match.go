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
