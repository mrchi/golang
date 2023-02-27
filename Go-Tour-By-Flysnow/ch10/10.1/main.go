package main

import (
	"fmt"
	"sync"
	"time"
)

func watchDog(stopChan <-chan bool, name string) {
	// for 循环保证一直在进行 select
	for {
		// 用 select + channel 做检测退出 goroutine
		select {
		case <-stopChan:
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

	stopChan := make(chan bool)

	go func() {
		defer wg.Done()
		watchDog(stopChan, "WatchDog001")
	}()

	time.Sleep(3 * time.Second)

	// 通过 channel 发停止指令给 goroutine
	stopChan <- true

	wg.Wait()
}
