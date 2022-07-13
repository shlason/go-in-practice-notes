package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	until := time.After(260 * time.Millisecond)

	go send(ch)

	for {
		select {
		case <-ch:
			fmt.Println("Got message")
		case <-until:
			fmt.Println("Time out")
			return
		default:
			fmt.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan<- bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	fmt.Println("Sent and closed")
}
