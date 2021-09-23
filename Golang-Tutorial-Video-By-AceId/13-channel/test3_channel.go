// 关闭 channel
package main

import (
	"fmt"
	"time"
)

func main() {
	// 向 nil channel 中发送数据会阻塞，因此最好都用 make 定义 channel
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// 关闭 channel
		// 关闭后不能再写数据，但可以收数据
		close(ch)
		fmt.Println("goroutine ended, channel closed.")
	}()

	time.Sleep(2 * time.Second)

	for {
		// 取 channel 中数据，ok 为是否成功的标识
		if num, ok := <-ch; ok {
			fmt.Printf("num = %v\n", num)
		} else {
			break
		}
	}

	fmt.Println("Finished.")
}
