//go:build !monadic_production

package main

import (
	"fmt"

	. "github.com/arcane-craft/monadic/exception"
	. "github.com/arcane-craft/monadic/io"
	. "github.com/arcane-craft/monadic/monad"
)

func MonadicDoCase() {
	fmt.Println("run monadic case with do syntax...")
	Init := LiftF1R(Init)
	Connect := LiftF1P1R(Connect)
	Send := LiftF2P(Send)
	Recv := LiftF1P1R(Recv)
	Println := LiftFVarP1R(fmt.Println)

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
