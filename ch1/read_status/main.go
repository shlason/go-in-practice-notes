package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Connects over TCP
	conn, _ := net.Dial("tcp", "golang.org:80")
	// Sends string over the connection
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// Print the first response line
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
