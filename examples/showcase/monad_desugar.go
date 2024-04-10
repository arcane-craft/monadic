package main

import (
	"fmt"

	. "github.com/arcane-craft/monadic/exception"
	. "github.com/arcane-craft/monadic/io"
	. "github.com/arcane-craft/monadic/monad"

	ikjxv1ef "github.com/arcane-craft/monadic"
)

func MonadicCase() {
	fmt.Println("run monadic case ...")
	Init := FFI1R(Init)
	Connect := FFI1P1R(Connect)
	Send := FFI2P(Send)
	Recv := FFI1P1R(Recv)
	Println := FFIVarP1R(fmt.Println)

	Perform(
		Catch(
			Then(ikjxv1ef.Zero[IO[int]](), func() IO[int] {
				return Bind(Init(), func(conf *Config) IO[int] {
					return Bind(Connect(conf.Remote), func(conn *Connection) IO[int] {
						return Then(Descript(Send(conn, []byte("hello")), "sending 'hello'"), func() IO[int] {
							return Bind(Recv(conn), func(ack []byte) IO[int] {
								return Println(string(ack))
							})
						})
					})
				})
			}),
			func(err error) IO[int] {
				return Println(err)
			},
		),
	)
}
