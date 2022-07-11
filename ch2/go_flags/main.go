package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish language."`
}

func main() {
	flags.Parse(&opts)

	if opts.Spanish {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
