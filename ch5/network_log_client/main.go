package main

import (
	"fmt"
	"log"
	"net"
)

const address string = "localhost:1902"

// Terminal 1
// $ nc -lk 1902

// Terminal 2
// $ go run main.go
func main() {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to %s", address))
	}
	defer func() {
		fmt.Println("Deferred function")
		conn.Close()
	}()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
}
