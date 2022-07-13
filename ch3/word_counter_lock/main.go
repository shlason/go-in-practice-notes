package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// In current directory and run command below
// go run main.go *.txt
func main() {
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			if err := tallyWords(filename, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once:")
	for word, count := range w.found {
		fmt.Printf("%s: %v\n", word, count)
	}
}

type words struct {
	sync.Mutex
	found map[string]int
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	conunt, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = conunt + n
}

func newWords() *words {
	return &words{found: make(map[string]int)}
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
