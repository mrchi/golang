package main

import (
	"fmt"
	"log"
	"net/http"
)

const listenAddr string = "localhost:8000"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s\n", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/", hello)

	log.Printf("Listening on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, nil)
}
