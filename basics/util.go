package basics

func Zero[A any]() A {
	var zero A
	return zero
}

func Id[A any](a A) A {
	return a
}

func Const[A, B any](a A, _ B) A {
	return a
}
