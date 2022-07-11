package main

import (
	"flag"
	"fmt"
)

const spanishFlagDesc = "Use Spanish language."

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "s", false, spanishFlagDesc)
	flag.BoolVar(&spanish, "spanish", false, spanishFlagDesc)
}

func main() {
	flag.Parse()
	if spanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
