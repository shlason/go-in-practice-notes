package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const address string = "localhost:1902"

// Terminal 1
// $ nc -luk 1902

// Terminal 2
// $ go run main.go

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", address, timeout)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to %s", address))
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
