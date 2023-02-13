package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 1; i <= delay; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, "sleeping", i, "seconds")
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("send", 2)
	fmt.Println("*** sending value ***")
	myChannel <- "a"
	fmt.Println("*** sending value ***")
	myChannel <- "b"
}

// 发送操作阻塞发送goroutine，直到另一个goroutine在同一channel上执行了接收操作。
// 反之亦然：接收操作阻塞接收goroutine，直到另一个goroutine在同一channel上执行了发送操作
func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("main", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
