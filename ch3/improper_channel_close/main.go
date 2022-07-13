package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	go send(msg)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			close(msg)
			// Pause to ensure see the failure before the main goroutine exits
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan<- string) {
	for {
		ch <- "Hello"
		time.Sleep(500 * time.Millisecond)
	}
}
