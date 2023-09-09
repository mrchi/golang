package main

import (
	"fmt"
	"log"
	"net/http"
)

const LISTENING_ADDR string = "localhost:8000"

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/a":
		fmt.Fprintln(w, "Serve /a")
	case "/b":
		fmt.Fprintln(w, "Serve /b")
	case "/c":
		fmt.Fprintln(w, "Serve /c")
	default:
		http.Error(w, "404 NOT FOUND", 404)
	}
}

func main() {
	log.Printf("Listening on %s\n", LISTENING_ADDR)
	http.ListenAndServe(LISTENING_ADDR, myHandler{})
}
