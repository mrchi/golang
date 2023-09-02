package main

import (
	"fmt"
	"log"
	"net"
)

const (
	TARGET        string = "192.168.50.1"
	WORKERS_COUNT int    = 100
)

func main() {
	ports := make(chan int)
	results := make(chan int)
	var openPorts []int

	// 创建 100 个 goroutine worker
	for i := 0; i < WORKERS_COUNT; i++ {
		go worker(ports, results)
	}

	// 向 ports 中写入端口
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// 从 results 中读取开放端口
	for i := 1; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	// 输出结果
	fmt.Println(openPorts)
}

func worker(ports, results chan int) {
	for p := range ports {
		// debug 日志
		log.Printf("worker: scanning port %d\n", p)
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", TARGET, p))
		if err != nil {
			results <- 0
			continue
		} else {
			conn.Close()
			results <- p
		}
	}
}
