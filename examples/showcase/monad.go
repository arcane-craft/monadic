package main

import (
	"fmt"

	. "github.com/arcane-craft/monadic/exception"
	. "github.com/arcane-craft/monadic/io"
	. "github.com/arcane-craft/monadic/monad"
)

func MonadicDoCase() {
	fmt.Println("run monadic case with do syntax...")
	Init := FFI1R(Init)
	Connect := FFI1P1R(Connect)
	Send := FFI2P(Send)
	Recv := FFI1P1R(Recv)
	Println := FFIVarP1R(fmt.Println)

	Perform(
		Catch(
			Do(func() IO[int] {
				conf := Init().X()
				conn := Connect(conf.Remote).X()
				Descript(Send(conn, []byte("hello")), "sending 'hello'").X()
				ack := Recv(conn).X()
				return Println(string(ack))
			}),
			func(err error) IO[int] {
				return Println(err)
			},
		),
	)
}
