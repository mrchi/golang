package main

import (
	"io"
	"log"
	"net"
)

const listenAddr string = "localhost:8000"

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Listening on %s", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Accepted connection from %s", conn.RemoteAddr())
		go echo(conn)
	}
}
