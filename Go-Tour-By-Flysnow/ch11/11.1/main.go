package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// for select循环有两种模式，一种是上一章的无限循环模式，只有收到终止指令才会退出
// 第二种模式是for range select有限循环，一般用于把可以迭代的内容发送到channel上
func generateNums(ctx context.Context, out chan<- int) {

	for _, i := range []int{1, 2, 3, 4, 5} {
		select {
		case <-ctx.Done():
			fmt.Println("End.")
			return
		case out <- i:
			fmt.Printf("Write %d to out.\n", i)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	out := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		generateNums(ctx, out)
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-out)
	}

	stop()
	wg.Wait()
}
