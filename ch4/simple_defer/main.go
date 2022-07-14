package main

import "fmt"

func main() {
	defer goodbye()

	fmt.Println("Hello")
}

func goodbye() {
	fmt.Println("goodbye")
}
