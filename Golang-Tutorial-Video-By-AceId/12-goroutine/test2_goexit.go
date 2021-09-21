// exit goroutine
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 定义匿名函数，并执行
	go func() {
		defer fmt.Println("A1.defer")

		// 定义带参数匿名函数，并传参执行，接收返回值
		a := func(b int) int {
			defer fmt.Println("B.defer")

			// 直接退出到当前 goroutine 的顶级调用，已经入栈过的 defer 还是要执行的
			runtime.Goexit()

			fmt.Printf("B = %v\n", b)
			return b + 1
		}(3)

		fmt.Printf("A = %v\n", a)
	}()

	time.Sleep(1 * time.Second)
}
