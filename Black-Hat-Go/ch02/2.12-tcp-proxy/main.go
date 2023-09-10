package main

import (
	"io"
	"log"
	"net"
)

const (
	listenAddr          string = "localhost:80"
	DESTINATION_ADDRESS string = "httpbin.org:80"
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listening on", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	dst, err := net.Dial("tcp", DESTINATION_ADDRESS)
	if err != nil {
		log.Fatalln(err)
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, conn); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(conn, dst); err != nil {
		log.Fatalln(err)
	}
}
