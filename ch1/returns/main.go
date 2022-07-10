package main

import (
	"fmt"
)

func Names() (string, string) {
	return "foo", "bar"
}

func main() {
	s1, s2 := Names()
	fmt.Println(s1, s2)

	// Skip second
	s3, _ := Names()
	fmt.Println(s3)
}
