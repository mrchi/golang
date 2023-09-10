package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

const listenAddr string = "localhost:8000"

var (
	hostProxy = make(map[string]string)
	proxies   = make(map[string]*httputil.ReverseProxy)
)

func init() {
	hostProxy["t1.example.com"] = "https://httpbin.org"
	hostProxy["t2.example.com"] = "http://httpbin.org"

	for k, v := range hostProxy {
		remote, err := url.Parse(v)
		if err != nil {
			log.Fatalf("Unable to parse proxy target %s", v)
		}
		proxies[k] = httputil.NewSingleHostReverseProxy(remote)
	}
}

func main() {
	r := mux.NewRouter()
	for host, proxy := range proxies {
		r.Host(host).Handler(proxy)
	}
	log.Printf("Listening on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, r)
}
