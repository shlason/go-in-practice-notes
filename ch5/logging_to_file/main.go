package main

import (
	"log"
	"os"
)

func main() {
	logFile, _ := os.Create("./log.txt")
	defer logFile.Close()

	logger := log.New(logFile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
	// Never execute
	logger.Println("End of th function.")
}
