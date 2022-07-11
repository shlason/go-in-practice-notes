package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("conf.yaml")
	defer file.Close()

	conf := Configuration{}
	b, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatalf("(ReadFile) error: %v", err)
	}
	err = yaml.Unmarshal([]byte(b), &conf)
	if err != nil {
		log.Fatalf("(Yaml Unmarshal) error: %v", err)
	}
	fmt.Print(conf.Enabled, conf.Path)
}
