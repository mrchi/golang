package main

import (
	"fmt"
	"log"
	"net/http"
)

const LISTEN_ADDR string = "localhost:8000"

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Start")
	l.Inner.ServeHTTP(w, r)
	log.Println("Finished")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	log.Printf("Listening on %s\n", LISTEN_ADDR)
	http.ListenAndServe(LISTEN_ADDR, &l)
}
