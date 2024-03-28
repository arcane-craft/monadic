package lazy

type (
	Byte = Value[byte]
	Rune = Value[rune]

	String = Value[string]
	Bytes  = Value[[]byte]
	Runes  = Value[[]rune]
)

func ToChars[B ~[]byte | ~[]rune, A ~string](v Value[A]) Value[B] {
	return New(func() B {
		return B(Eval(v))
	})
}

func ToString[B ~string, A ~string | ~[]byte | ~[]rune](v Value[A]) Value[B] {
	return New(func() B {
		return B(Eval(v))
	})
}

func ToSlice[B, A ~[]E, E any](a Value[A]) Value[B] {
	return New(func() B {
		return B(Eval(a))
	})
}
