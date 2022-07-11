package main

import (
	"fmt"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "The homepage.")
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

// Antipattern
func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:8080", nil)
}
