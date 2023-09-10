package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const LISTENING_ADDR string = "localhost:8000"

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static-files")))

	log.Printf("Listening on %s\n", LISTENING_ADDR)
	http.ListenAndServe(LISTENING_ADDR, r)

}
