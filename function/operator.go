package function

func Compose[A, B, C any](l func(B) C, r func(A) B) func(A) C {
	return func(a A) C {
		return l(r(a))
	}
}
