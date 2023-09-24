package main

import (
	"fmt"
	"log"
	"net/http"
)

const LISTEN_ADDR string = "localhost:8000"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s\n", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/", hello)

	log.Printf("Listening on %s\n", LISTEN_ADDR)
	http.ListenAndServe(LISTEN_ADDR, nil)
}
