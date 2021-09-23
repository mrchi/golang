// channel
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个 channel，channel 用于在 goroutine 之间传递数据
	ch := make(chan int)

	go func() {
		defer fmt.Println("goroutine ended.")

		fmt.Println("goroutine running...")

		ch <- 123 // 将 123 写入 channel
	}()

	for i := 0; i < 3; i++ {
		fmt.Printf("Sleep %v second(s)...\n", i)
		time.Sleep(1 * time.Second)
	}

	a := <-ch // 将 channel 中的值赋给变量

	fmt.Printf("a = %v\n", a)
	fmt.Println("main ended.")

	// 无缓冲 channel，两个 goroutine 的代码没有同时执行到管道操作语句时，先到的 goroutine 会阻塞等待
}
