package function

func Curry[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

func Uncurry[A, B, C any](f func(A) func(B) C) func(A, B) C {
	return func(a A, b B) C {
		return f(a)(b)
	}
}

func Curry3[A, B, C, D any](f func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(B) func(C) D {
		return func(b B) func(C) D {
			return func(c C) D {
				return f(a, b, c)
			}
		}
	}
}

func Uncurry3[A, B, C, D any](f func(A) func(B) func(C) D) func(A, B, C) D {
	return func(a A, b B, c C) D {
		return f(a)(b)(c)
	}
}

func Curry4[A, B, C, D, E any](f func(A, B, C, D) E) func(A) func(B) func(C) func(D) E {
	return func(a A) func(B) func(C) func(D) E {
		return func(b B) func(C) func(D) E {
			return func(c C) func(D) E {
				return func(d D) E {
					return f(a, b, c, d)
				}
			}
		}
	}
}

func Uncurry4[A, B, C, D, E any](f func(A) func(B) func(C) func(D) E) func(A, B, C, D) E {
	return func(a A, b B, c C, d D) E {
		return f(a)(b)(c)(d)
	}
}

func Curry5[A, B, C, D, E, F any](f func(A, B, C, D, E) F) func(A) func(B) func(C) func(D) func(E) F {
	return func(a A) func(B) func(C) func(D) func(E) F {
		return func(b B) func(C) func(D) func(E) F {
			return func(c C) func(D) func(E) F {
				return func(d D) func(E) F {
					return func(e E) F {
						return f(a, b, c, d, e)
					}
				}
			}
		}
	}
}

func Uncurry5[A, B, C, D, E, F any](f func(A) func(B) func(C) func(D) func(E) F) func(A, B, C, D, E) F {
	return func(a A, b B, c C, d D, e E) F {
		return f(a)(b)(c)(d)(e)
	}
}
