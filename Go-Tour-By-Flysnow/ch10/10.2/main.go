package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func watchDog(ctx context.Context, name string) {
	// for 循环保证一直在进行 select
	for {
		select {
		// Done方法返回一个只读的channel，类型为struct{}。
		// 在协程中，如果该方法返回的chan可以读取，则意味着Context已经发起了取消信号。
		// 通过Done方法收到这个信号后，就可以做清理操作，然后退出协程，释放资源。
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
	wg.Add(1)

	// Context是一个接口，
	// 它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。
	// 一旦取消指令下达，那么被Context跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。

	// Go语言提供了可以帮助我们生成不同Context的函数，通过这些函数可以生成一棵Context树。
	// 这样Context才可以关联起来，父Context发出取消信号的时候，子Context也会发出，这样就可以控制不同层级协程的退出。

	// 通过context.Background()获取一个根节点Context，是空Context
	// WithCancel(parent Context)：生成一个可取消的Context。
	ctx, stop := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		watchDog(ctx, "WatchDog001")
	}()

	time.Sleep(3 * time.Second)

	// 调用取消函数
	stop()

	wg.Wait()
}
