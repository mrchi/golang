package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const listenAddr string = "localhost:8000"

type badAuth struct {
	Username string
	Password string
}

func (ba *badAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if username != ba.Username || password != ba.Password {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(ctx)

	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w, "Hello, %s", username)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")

	n := negroni.Classic()
	n.Use(&badAuth{Username: "foo", Password: "bar"})
	n.UseHandler(r)

	log.Printf("Listening on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, n)
}
