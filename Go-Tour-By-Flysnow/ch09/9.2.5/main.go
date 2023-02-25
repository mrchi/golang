package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 而sync.Cond可以用于发号施令，一声令下所有协程都可以开始执行，
	// 关键点在于协程开始的时候是等待状态，要等待sync.Cond唤醒才能执行。

	// 通过sync.NewCond函数生成一个*sync.Cond，用于阻塞和唤醒协程。
	cond := sync.NewCond(&sync.Mutex{})

	var wg sync.WaitGroup
	wg.Add(11)

	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Printf("Engine %d is ready.\n", num)

			// 使用的时候需要加锁，使用sync.Cond中的锁即可，也就是L字段。
			cond.L.Lock()
			// Wait，阻塞当前协程，直到被其他协程通过调用Broadcast或者Signal方法唤醒。
			cond.Wait()

			fmt.Printf("Engine %d fired.\n", num)

			cond.L.Unlock()
		}(i)
		time.Sleep(500 * time.Millisecond)
	}

	// 在调用Signal或者Broadcast之前，要确保目标协程处于Wait阻塞状态，不然会出现死锁问题
	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()
		fmt.Println("Launch ready.")
		fmt.Println()

		// Signal，唤醒一个等待时间最长的协程。
		fmt.Println("Fire earliest engine!")
		cond.Signal()

		time.Sleep(1 * time.Second)
		fmt.Println()

		// Broadcast，唤醒所有等待的协程。
		fmt.Println("Fire all other engines!")
		cond.Broadcast()
	}()

	wg.Wait()
}
