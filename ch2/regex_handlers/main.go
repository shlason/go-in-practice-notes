package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goodbye)

	http.ListenAndServe(":8080", pr)
}

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func (rr *regexResolver) Add(pattern string, handlerFunc http.HandlerFunc) {
	rr.handlers[pattern] = handlerFunc
	cache, _ := regexp.Compile(pattern)
	rr.cache[pattern] = cache
}

func (rr *regexResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range rr.handlers {
		if rr.cache[pattern].MatchString(check) {
			handlerFunc(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
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
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Empty"
	}
	fmt.Fprint(w, "Goodbye ", name)
}
