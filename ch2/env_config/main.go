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

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println(os.Getenv("PORT"))
	http.ListenAndServe("localhost:"+os.Getenv("PORT"), nil)
}
