package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to open port 8080")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket")
		conn.Close()
	}

	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	conn.Write(data)
	panic("Failure in response!")
}
