package monadic

type Unit struct{}

type Data[A any, _T any] interface {
	Kind() _T
}

type Generalize[D Data[A, _T], A any, _T any] interface {
	Concretize(Data[any, _T]) D
	Abstract(D) Data[any, _T]
}
