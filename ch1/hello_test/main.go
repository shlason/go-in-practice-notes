package main

import "fmt"

func getName() string {
	return "Jason"
}

func main() {
	name := getName()
	fmt.Printf("Hello %s\n", name)
}
