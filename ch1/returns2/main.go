package main

import (
	"fmt"
)

func Names() (first string, second string) {
	first = "foo"
	second = "bar"
	return
}

func main() {
	s1, s2 := Names()
	fmt.Println(s1, s2)
}
