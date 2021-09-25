// channel 和 range
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("goroutine ended, channel closed.")
	}()

	time.Sleep(2 * time.Second)

	// 使用 range 不断从 channel 中取数据
	for num := range ch {
		fmt.Printf("num = %v\n", num)
	}

	fmt.Println("Finished.")
}
