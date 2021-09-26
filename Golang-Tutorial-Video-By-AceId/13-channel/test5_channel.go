// channel 和 select
package main

import "fmt"

func fibonacci(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		// 使用 select 来同时处理多个 channel 的情况
		select {
		// 如果 channel c 可写
		case c <- x:
			x, y = y, x+y
		// 如果 channel quit 可读
		case <-quit:
			fmt.Println("Fabonacci quited.")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
