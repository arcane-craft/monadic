package main

import (
	"fmt"

	"github.com/arcane-craft/monadic"
	"github.com/arcane-craft/monadic/either"
	"github.com/arcane-craft/monadic/io"
	"github.com/arcane-craft/monadic/lazy"
)

type Config struct {
	Remote string
}

type Connection struct{}

func Init() (*Config, error) {
	fmt.Println("Init()")
	// return &Config{
	// 	Remote: "localhost:80",
	// }, nil
	return nil, fmt.Errorf("config not found")
}

func Connect(remote string) (*Connection, error) {
	fmt.Println("Connect()", remote)
	return &Connection{}, nil
	// return nil, fmt.Errorf("unreachable remote")
}

func Send(conn *Connection, msg []byte) error {
	fmt.Println("Send()", conn, string(msg))
	return nil
	// return fmt.Errorf("connection reset")
}

func Recv(conn *Connection) ([]byte, error) {
	fmt.Println("Recv()", conn)
	return []byte("world!"), nil
}

func main() {
	func() {
		fmt.Println("run classic case...")
		conf, err := Init()
		if err != nil {
			conf = &Config{Remote: "localhost:443"}
		}
		conn, err := Connect(conf.Remote)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		err = Send(conn, []byte("hello"))
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		ack, err := Recv(conn)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(string(ack))
	}()
	// =>
	func() {
		fmt.Println("run monadic case...")
		Init := io.Lift1R(Init)
		Connect := io.Lift1P1R(Connect)
		Send := io.Lift2P(Send)
		Recv := io.Lift1P1R(Recv)
		content, err := io.Eval(
			io.Bind(
				io.Try(Init()),
				func(conf lazy.Value[either.Either[*Config, error]]) io.IO[[]byte] {
					return io.Bind(
						Connect(either.EitherOf(
							conf,
							lazy.Lift(func(c *Config) string { return c.Remote }),
							lazy.Lift(func(_ error) string { return "localhost:443" }),
						)),
						func(conn lazy.Value[*Connection]) io.IO[[]byte] {
							return io.Bind(
								Send(conn, lazy.Const([]byte("hello"))),
								func(lazy.Value[monadic.Void]) io.IO[[]byte] {
									return io.Bind(
										Recv(conn),
										func(ack lazy.Value[[]byte]) io.IO[[]byte] {
											return io.Ret(ack)
										},
									)
								},
							)
						},
					)
				},
			),
		)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		fmt.Println(string(content))
	}()
	// =>
	func() {
		fmt.Println("run monadic case by do syntax...")
		content, err := io.Eval(
			io.Do(func(ctx *io.Context[[]byte]) io.IO[[]byte] {
				Init := io.Lift1R(Init)
				ConnectX := io.Lift1P1RX[[]byte](Connect)
				SendX := io.Lift2PX[[]byte](Send)
				RecvX := io.Lift1P1RX[[]byte](Recv)

				conf := io.From(ctx, io.Try(Init()))
				remote := either.EitherOf(
					conf,
					lazy.Lift(func(c *Config) string { return c.Remote }),
					lazy.Lift(func(_ error) string { return "localhost:443" }),
				)
				conn := ConnectX(ctx, remote)
				SendX(ctx, conn, lazy.Const([]byte("hello")))
				ack := RecvX(ctx, conn)
				return io.Ret(ack)
			}),
		)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(string(content))
	}()
}
