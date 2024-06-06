package basics

func Zero[A any]() A {
	var zero A
	return zero
}

func Id[A any](a A) A {
	return a
}

func Const[B, A any](a A, _ B) A {
	return a
}
