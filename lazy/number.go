package lazy

type (
	Int   = Value[int]
	Int8  = Value[int8]
	Int16 = Value[int16]
	Int32 = Value[int32]
	Int64 = Value[int64]

	Uint   = Value[uint]
	Uint8  = Value[uint8]
	Uint16 = Value[uint16]
	Uint32 = Value[uint32]
	Uint64 = Value[uint64]

	Float32 = Value[float32]
	Float64 = Value[float64]
)

func ToNumber[B, A Number](a Value[A]) Value[B] {
	return New(func() B {
		return B(Eval(a))
	})
}
