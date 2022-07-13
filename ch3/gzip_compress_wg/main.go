package main

import (
	"compress/gzip"
	"io"
	"os"
	"sync"
)

// In current directory and run command below
// go run main.go ./exampledatas/*
func main() {
	var wg sync.WaitGroup
	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
		compress(file)
	}
	wg.Wait()
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
