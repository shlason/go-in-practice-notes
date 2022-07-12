package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/braintree/manners"
)

func main() {
	handler := newHandler()

	ch := make(chan os.Signal, 2)
	/**
	ref: https://github.com/braintree/manners/issues/45

	Regarding the documented signals example, os.Kill can't be caught, it's just there for killing other processes.
	os.Kill is like kill -9 which kills a process immediately
	syscall.SIGTERM is equivalent to kill which allows the process time to cleanup
	*/
	// signal.Notify(ch, os.Interrupt, os.Kill)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go listerForShutdown(ch)

	manners.ListenAndServe(":8080", handler)
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Empty"
	}
	fmt.Fprint(w, "Hello, my name is ", name)
}

func newHandler() *handler {
	return &handler{}
}

func listerForShutdown(c <-chan os.Signal) {
	<-c
	// Do somthing
	time.Sleep(3 * time.Second)
	manners.Close()
}
