package main

import (
	"io"
	"log"
	"net"
)

const LISTENING_ADDRESS string = "localhost:8000"

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", LISTENING_ADDRESS)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Listening on %s", LISTENING_ADDRESS)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Accepted connection from %s", conn.RemoteAddr())
		go echo(conn)
	}
}
