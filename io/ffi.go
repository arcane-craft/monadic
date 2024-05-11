package io

import (
	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/lazy"
	"github.com/arcane-craft/monadic/result"
	"github.com/arcane-craft/monadic/tuple"
)

func Perform[A any](m IO[A]) (A, error) {
	r := m.v()
	zero := monadic.Zero[A]()
	if result.IsOk(r) {
		return result.FromOk(zero, r), nil
	}
	return zero, result.FromFail(nil, r)
}

func FFISucc(f func()) func() IO[monadic.Void] {
	return func() IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f()
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI1PSucc[P any](f func(P)) func(P) IO[monadic.Void] {
	return func(p P) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(p)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI2PSucc[P1, P2 any](f func(P1, P2)) func(P1, P2) IO[monadic.Void] {
	return func(p1 P1, p2 P2) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(p1, p2)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI3PSucc[P1, P2, P3 any](f func(P1, P2, P3)) func(P1, P2, P3) IO[monadic.Void] {
	return func(p1 P1, p2 P2, p3 P3) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(p1, p2, p3)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI4PSucc[P1, P2, P3, P4 any](f func(P1, P2, P3, P4)) func(P1, P2, P3, P4) IO[monadic.Void] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(p1, p2, p3, p4)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI5PSucc[P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5)) func(P1, P2, P3, P4, P5) IO[monadic.Void] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(p1, p2, p3, p4, p5)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFIVarPSucc[P any](f func(...P)) func(...P) IO[monadic.Void] {
	return func(ps ...P) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(ps...)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI1PVarPSucc[P, VP any](f func(P, ...VP)) func(P, ...VP) IO[monadic.Void] {
	return func(p P, v ...VP) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(p, v...)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI2PVarPSucc[P1, P2, VP any](f func(P1, P2, ...VP)) func(P1, P2, ...VP) IO[monadic.Void] {
	return func(p1 P1, p2 P2, v ...VP) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				f(p1, p2, v...)
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI(f func() error) func() IO[monadic.Void] {
	return func() IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI1P[P any](f func(P) error) func(P) IO[monadic.Void] {
	return func(p P) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(p); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI2P[P1, P2 any](f func(P1, P2) error) func(P1, P2) IO[monadic.Void] {
	return func(p1 P1, p2 P2) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(p1, p2); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI3P[P1, P2, P3 any](f func(P1, P2, P3) error) func(P1, P2, P3) IO[monadic.Void] {
	return func(p1 P1, p2 P2, p3 P3) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(p1, p2, p3); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI4P[P1, P2, P3, P4 any](f func(P1, P2, P3, P4) error) func(P1, P2, P3, P4) IO[monadic.Void] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(p1, p2, p3, p4); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI5P[P1, P2, P3, P4, P5 any](f func(P1, P2, P3, P4, P5) error) func(P1, P2, P3, P4, P5) IO[monadic.Void] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(p1, p2, p3, p4, p5); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFIVarP[P any](f func(...P) error) func(...P) IO[monadic.Void] {
	return func(ps ...P) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(ps...); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI1PVarP[P, VP any](f func(P, ...VP) error) func(P, ...VP) IO[monadic.Void] {
	return func(p P, ps ...VP) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(p, ps...); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI2PVarP[P1, P2, VP any](f func(P1, P2, ...VP) error) func(P1, P2, ...VP) IO[monadic.Void] {
	return func(p1 P1, p2 P2, ps ...VP) IO[monadic.Void] {
		return IO[monadic.Void]{
			v: lazy.New(func() either.Either[error, monadic.Void] {
				if err := f(p1, p2, ps...); err != nil {
					return result.Fail[monadic.Void](err)
				}
				return result.Ok(monadic.Void{})
			}),
		}
	}
}

func FFI1R[R any](f func() (R, error)) func() IO[R] {
	return func() IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f()
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI1P1R[P, R any](f func(P) (R, error)) func(P) IO[R] {
	return func(p P) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(p)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI2P1R[P1, P2, R any](f func(P1, P2) (R, error)) func(P1, P2) IO[R] {
	return func(p1 P1, p2 P2) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(p1, p2)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI3P1R[P1, P2, P3, R any](f func(P1, P2, P3) (R, error)) func(P1, P2, P3) IO[R] {
	return func(p1 P1, p2 P2, p3 P3) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(p1, p2, p3)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI4P1R[P1, P2, P3, P4, R any](f func(P1, P2, P3, P4) (R, error)) func(P1, P2, P3, P4) IO[R] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(p1, p2, p3, p4)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI5P1R[P1, P2, P3, P4, P5, R any](f func(P1, P2, P3, P4, P5) (R, error)) func(P1, P2, P3, P4, P5) IO[R] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(p1, p2, p3, p4, p5)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFIVarP1R[P, R any](f func(...P) (R, error)) func(...P) IO[R] {
	return func(ps ...P) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(ps...)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI1PVarP1R[P, VP, R any](f func(P, ...VP) (R, error)) func(P, ...VP) IO[R] {
	return func(p P, ps ...VP) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(p, ps...)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI2PVarP1R[P1, P2, VP, R any](f func(P1, P2, ...VP) (R, error)) func(P1, P2, ...VP) IO[R] {
	return func(p1 P1, p2 P2, ps ...VP) IO[R] {
		return IO[R]{
			v: lazy.New(func() either.Either[error, R] {
				r, err := f(p1, p2, ps...)
				if err != nil {
					return result.Fail[R](err)
				}
				return result.Ok(r)
			}),
		}
	}
}

func FFI2R[R1, R2 any](f func() (R1, R2, error)) func() IO[tuple.Tuple2[R1, R2]] {
	return func() IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f()
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFI1P2R[P, R1, R2 any](f func(P) (R1, R2, error)) func(P) IO[tuple.Tuple2[R1, R2]] {
	return func(p P) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(p)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFI2P2R[P1, P2, R1, R2 any](f func(P1, P2) (R1, R2, error)) func(P1, P2) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 P1, p2 P2) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(p1, p2)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFI3P2R[P1, P2, P3, R1, R2 any](f func(P1, P2, P3) (R1, R2, error)) func(P1, P2, P3) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 P1, p2 P2, p3 P3) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(p1, p2, p3)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFI4P2R[P1, P2, P3, P4, R1, R2 any](f func(P1, P2, P3, P4) (R1, R2, error)) func(P1, P2, P3, P4) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(p1, p2, p3, p4)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFI5P2R[P1, P2, P3, P4, P5, R1, R2 any](f func(P1, P2, P3, P4, P5) (R1, R2, error)) func(P1, P2, P3, P4, P5) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(p1, p2, p3, p4, p5)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFIVarP2R[P, R1, R2 any](f func(...P) (R1, R2, error)) func(...P) IO[tuple.Tuple2[R1, R2]] {
	return func(ps ...P) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(ps...)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFI1PVarP2R[P, VP, R1, R2 any](f func(P, ...VP) (R1, R2, error)) func(P, ...VP) IO[tuple.Tuple2[R1, R2]] {
	return func(p P, ps ...VP) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(p, ps...)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}

func FFI2PVarP2R[P1, P2, VP, R1, R2 any](f func(P1, P2, ...VP) (R1, R2, error)) func(P1, P2, ...VP) IO[tuple.Tuple2[R1, R2]] {
	return func(p1 P1, p2 P2, ps ...VP) IO[tuple.Tuple2[R1, R2]] {
		return IO[tuple.Tuple2[R1, R2]]{
			v: lazy.New(func() either.Either[error, tuple.Tuple2[R1, R2]] {
				r1, r2, err := f(p1, p2, ps...)
				if err != nil {
					return result.Fail[tuple.Tuple2[R1, R2]](err)
				}
				return result.Ok(tuple.T2(r1, r2))
			}),
		}
	}
}
