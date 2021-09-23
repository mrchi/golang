// 有缓冲 channel
package main

import (
	"fmt"
	"time"
)

func main() {
	// 新建一个容量为 3 的有缓冲 channel
	ch := make(chan int, 3)

	go func() {
		defer println("goroutine ended.")
		for i := 0; i < 4; i++ {
			// 有缓冲 channel 不会阻塞，除非 channel 写满
			ch <- i
			fmt.Printf("Write %v to channel, length = %d, cap = %d\n", i, len(ch), cap(ch))
		}
	}()

	// 当 channel 写满时，goroutine 被阻塞， 等待 channel 有空间后再继续执行
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Printf("num = %v\n", num)
		time.Sleep(1 * time.Second)
	}
}
