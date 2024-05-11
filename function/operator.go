package function

func Compose[A, B, C any](l func(B) C, r func(A) B) func(A) C {
	return func(a A) C {
		return l(r(a))
	}
}

func Flip[A, B, C any](f func(A, B) C) func(B, A) C {
	return func(b B, a A) C {
		return f(a, b)
	}
}

func Infix[A, B, R any](a A, op func(A, B) R, b B) R {
	return op(a, b)
}

func Infix3[A, B, R1, C, R2 any](a A, op1 func(A, B) R1, b B, op2 func(R1, C) R2, c C) R2 {
	return op2(op1(a, b), c)
}

func Infix4[A, B, R1, C, R2, D, R3 any](a A, op1 func(A, B) R1, b B, op2 func(R1, C) R2, c C, op3 func(R2, D) R3, d D) R3 {
	return op3(op2(op1(a, b), c), d)
}

func Infix5[A, B, R1, C, R2, D, R3, E, R4 any](a A, op1 func(A, B) R1, b B, op2 func(R1, C) R2, c C, op3 func(R2, D) R3, d D, op4 func(R3, E) R4, e E) R4 {
	return op4(op3(op2(op1(a, b), c), d), e)
}

func Infix6[A, B, R1, C, R2, D, R3, E, R4, F, R5 any](a A, op1 func(A, B) R1, b B, op2 func(R1, C) R2, c C, op3 func(R2, D) R3, d D, op4 func(R3, E) R4, e E, op5 func(R4, F) R5, f F) R5 {
	return op5(op4(op3(op2(op1(a, b), c), d), e), f)
}

func Infix7[A, B, R1, C, R2, D, R3, E, R4, F, R5, G, R6 any](a A, op1 func(A, B) R1, b B, op2 func(R1, C) R2, c C, op3 func(R2, D) R3, d D, op4 func(R3, E) R4, e E, op5 func(R4, F) R5, f F, op6 func(R5, G) R6, g G) R6 {
	return op6(op5(op4(op3(op2(op1(a, b), c), d), e), f), g)
}

func Infix8[A, B, R1, C, R2, D, R3, E, R4, F, R5, G, R6, H, R7 any](a A, op1 func(A, B) R1, b B, op2 func(R1, C) R2, c C, op3 func(R2, D) R3, d D, op4 func(R3, E) R4, e E, op5 func(R4, F) R5, f F, op6 func(R5, G) R6, g G, op7 func(R6, H) R7, h H) R7 {
	return op7(op6(op5(op4(op3(op2(op1(a, b), c), d), e), f), g), h)
}

func Infix9[A, B, R1, C, R2, D, R3, E, R4, F, R5, G, R6, H, R7, I, R8 any](a A, op1 func(A, B) R1, b B, op2 func(R1, C) R2, c C, op3 func(R2, D) R3, d D, op4 func(R3, E) R4, e E, op5 func(R4, F) R5, f F, op6 func(R5, G) R6, g G, op7 func(R6, H) R7, h H, op8 func(R7, I) R8, i I) R8 {
	return op8(op7(op6(op5(op4(op3(op2(op1(a, b), c), d), e), f), g), h), i)
}
