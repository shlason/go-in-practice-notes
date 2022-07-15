package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("Inside deferred function")
		if err := recover(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		fmt.Println("End deferred function")
	}()
	log.Println("This is a regular message.")
	log.Panicln("This is panic log.")
}
