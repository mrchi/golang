// goroutine
package main

import (
	"fmt"
	"time"
)

func newTask() {
	for i := 0; i < 10; i++ {
		fmt.Printf("New goroutine, i = %v\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个 goroutine，执行 newTask() 函数
	go newTask()

	for i := 0; i < 5; i++ {
		fmt.Printf("Main goroutine, i = %v\n", i)
		time.Sleep(1 * time.Second)
	}
	// 主进程结束之时，goroutine 也会终止（不管 goroutine 是否执行完成）
}
