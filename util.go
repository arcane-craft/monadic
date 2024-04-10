package monadic

func Zero[A any]() A {
	var zero A
	return zero
}
