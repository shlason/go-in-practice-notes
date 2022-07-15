package main

import "log"

func main() {
	log.Println("This is a regular message.")
	log.Fatalln("This is a fatal error.")
	// Never execute
	log.Println("This is the end of the function.")
}
