package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s stopped.\n", name)
			return
		default:
			fmt.Printf("%s is watching...\n", name)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ctx, stop := context.WithCancel(context.Background())

	for i := 1; i < 4; i++ {
		go func(id int) {
			defer wg.Done()
			watchDog(ctx, fmt.Sprintf("WatchDog00%d", id))
		}(i)
	}

	time.Sleep(3 * time.Second)

	// 调用取消函数，同时取消多个 goroutine
	stop()

	wg.Wait()
}
