package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type key string

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[getUser]Quit.")
			return
		default:
			userId := ctx.Value(key("userId"))
			fmt.Printf("[getUser]ID=%v\n", userId)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, stop := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, key("userId"), 2)

	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(3 * time.Second)

	stop()

	wg.Wait()
}
