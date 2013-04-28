package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(conn, "Oi pessoal do #devincachu!")
		conn.Close()
	}
}
