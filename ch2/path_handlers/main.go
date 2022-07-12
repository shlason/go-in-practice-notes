package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)

	http.ListenAndServe(":8080", pr)
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(w, r)
			return
		} else if err != nil {
			fmt.Fprint(w, err)
		}
	}
	http.NotFound(w, r)
}

func newPathResolver() *pathResolver {
	return &pathResolver{handlers: make(map[string]http.HandlerFunc)}
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
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Empty"
	}
	fmt.Fprint(w, "Goodbye ", name)
}
