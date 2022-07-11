package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

type Configuration struct {
	Section struct {
		Enabled bool
		Path    string
	}
}

func main() {
	config := Configuration{}

	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
