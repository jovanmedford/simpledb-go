package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1099")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("database server ready.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection:", conn.RemoteAddr())

	// Echo to client
	if _, err := io.Copy(conn, conn); err != nil {
		fmt.Println("Error echoing data:", err)
	}
}
