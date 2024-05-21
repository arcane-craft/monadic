//go:build monadic_production

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
			Then(DoInit[IO[int]](), func () IO[int] { return Bind(Init(), func (conf *Config) IO[int] { return Bind(Connect(conf.Remote), func (conn *Connection) IO[int] { return Then(Descript(Send(conn, []byte("hello")), "sending 'hello'"), func () IO[int] { return Bind(Recv(conn), func (ack []byte) IO[int] { return Println(string(ack)) }) }) }) }) }),
			func(err error) IO[int] {
				return Println(err)
			},
		),
	)
}
