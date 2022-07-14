package main

import "errors"

func main() {
	panic(errors.New("something bad happened"))
}
