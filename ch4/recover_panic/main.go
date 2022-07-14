package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
		fmt.Println("defer func")
	}()
	fmt.Println("Before oops")
	oops()
	fmt.Println("After oops")
}

func oops() {
	panic(errors.New("something bad happened"))
}
