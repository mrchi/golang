package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	result := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		time.Sleep(5 * time.Second)
		result <- "Get result"
	}()

	// select timeout模式的核心在于通过time.After函数设置一个超时时间，防止因为异常造成select语句的无限等待。
	// 如果可以使用Context的WithTimeout函数超时取消，则优先使用。
	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout.")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
