package main

import (
	"flag"
	"fmt"
)

const spanishFlagDesc = "Use Spanish language."

var (
	name    string
	spanish bool
)

func init() {
	flag.BoolVar(&spanish, "s", false, spanishFlagDesc)
	flag.BoolVar(&spanish, "spanish", false, spanishFlagDesc)
	flag.StringVar(&name, "name", "World", "A name to say hello to.")
}

func main() {
	flag.Parse()
	if spanish {
		fmt.Printf("Hola %s!\n", name)
	} else {
		fmt.Printf("Hello %s!\n", name)
	}
}
