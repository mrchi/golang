// openssl req -nodes -x509 -newkey rsa:4096 -keyout keys/serverKey.pem -out keys/serverCrt.pem -days 365 -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost"
// openssl req -nodes -x509 -newkey rsa:4096 -keyout keys/clientKey.pem -out keys/clientCrt.pem -days 365
// xhs --verify=no --cert=keys/clientCrt.pem --cert-key=keys/clientKey.pem  localhost:8000

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

const KEYS_PATH = "keys/"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello: %s\n", r.TLS.PeerCertificates[0].Subject.CommonName)
	fmt.Fprint(w, "Authentication successful")
}

func runServer(addr string) {
	http.HandleFunc("/", helloHandler)

	clientCert, err := os.ReadFile(path.Join(KEYS_PATH, "clientCrt.pem"))
	if err != nil {
		log.Fatalln(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(clientCert)

	tlsConf := &tls.Config{
		ClientCAs:  pool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	log.Printf("Listening on %s\n", addr)
	server := &http.Server{
		Addr:      addr,
		TLSConfig: tlsConf,
	}
	server.ListenAndServeTLS(
		path.Join(KEYS_PATH, "serverCrt.pem"),
		path.Join(KEYS_PATH, "serverKey.pem"),
	)
}

func runClient(target string) {
	cert, err := tls.LoadX509KeyPair(
		path.Join(KEYS_PATH, "clientCrt.pem"),
		path.Join(KEYS_PATH, "clientKey.pem"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	serverCert, err := os.ReadFile(path.Join(KEYS_PATH, "serverCrt.pem"))
	if err != nil {
		log.Fatalln(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(serverCert)

	tlsConf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConf,
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(target)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Success: %s\n", body)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: server/client")
	}
	switch os.Args[1] {
	case "server":
		runServer("localhost:8000")
	case "client":
		runClient("https://localhost:8000")
	default:
		log.Fatalln("Usage: server/client")
	}
}
