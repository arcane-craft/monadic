package monadic

type Void struct{}

type Nillable interface {
	IsNil() bool
}

type Data[A any, E Nillable] interface {
	Resolve() (A, E)
}
