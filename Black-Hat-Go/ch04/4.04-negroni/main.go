package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const LISTEN_ADDR string = "localhost:8000"

type trivial struct{}

func (t trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Executing trivial middleware")
	next(w, r)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc(
		"/users/{user:[a-zA-z]+}",
		func(w http.ResponseWriter, r *http.Request) {
			user := mux.Vars(r)["user"]
			fmt.Fprintf(w, "Hello, %s", user)
			log.Println("Executing handler function")
		},
	).Methods("GET").Host("foo.example.com")

	n := negroni.Classic()
	// Use 的先后顺序决定了中间件和 handler 的执行顺序
	n.Use(trivial{})
	n.UseHandler(r)

	log.Printf("Listening on %s\n", LISTEN_ADDR)
	http.ListenAndServe(LISTEN_ADDR, n)
}
