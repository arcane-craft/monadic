package main

import "fmt"

type Config struct {
	Remote string
}

type Connection struct{}

func Init() (*Config, error) {
	fmt.Println("Init()")
	return &Config{
		Remote: "localhost:80",
	}, nil
	// return nil, fmt.Errorf("config not found")
}

func Connect(remote string) (*Connection, error) {
	fmt.Println("Connect()", remote)
	return &Connection{}, nil
	// return nil, fmt.Errorf("unreachable remote")
}

func Send(conn *Connection, msg []byte) error {
	fmt.Println("Send()", conn, string(msg))
	// return nil
	return fmt.Errorf("connection reset")
}

func Recv(conn *Connection) ([]byte, error) {
	fmt.Println("Recv()", conn)
	return []byte("world!"), nil
}

func ClassicCase() {
	fmt.Println("run classic case...")
	conf, err := Init()
	if err != nil {
		conf = &Config{Remote: "localhost:443"}
	}
	conn, err := Connect(conf.Remote)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Send(conn, []byte("hello"))
	if err != nil {
		fmt.Println(fmt.Errorf("sending 'hello' failed: %w", err))
		return
	}
	ack, err := Recv(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(ack))
}
