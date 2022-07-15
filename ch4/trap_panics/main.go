package main

import (
	"errors"
	"fmt"
	"goipn/ch4/safely"
	"time"
)

func message() {
	fmt.Println("Inside goroutine")
	panic(errors.New("oops"))
}

func main() {
	safely.Go(message)
	fmt.Println("Outside goroutine")
	time.Sleep(1 * time.Second)
}
