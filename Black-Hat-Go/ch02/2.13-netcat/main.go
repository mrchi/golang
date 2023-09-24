package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

const LISTEN_ADDR string = "localhost:8000"

func main() {
	listener, err := net.Listen("tcp", LISTEN_ADDR)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listening on", LISTEN_ADDR)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Accept connection from %s", conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	command := exec.Command("/bin/sh", "-i")
	reader, writer := io.Pipe()
	command.Stdin = conn
	command.Stdout = writer
	go io.Copy(conn, reader)

	command.Run()
}
