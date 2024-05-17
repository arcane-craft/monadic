package num

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Num interface {
	Integer | Float
}

func Plus[A Num](a, b A) A {
	return a + b
}

func Mult[A Num](a, b A) A {
	return a * b
}
