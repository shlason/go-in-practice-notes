package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Empty"
	}
	fmt.Fprint(w, "Hello, my name is ", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	parts := strings.Split(path, "/")
	fmt.Println(parts[0] == "")
	name := parts[2]
	if name == "" {
		name = "Empty"
	}
	fmt.Fprint(w, "Goodbye ", name)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "The homepage.")
}
