package main

import (
	"fmt"
	"sync"
	"time"
)

var sum = 0
var mutex sync.RWMutex //读写锁sync.RWMutex

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func readSum() int {
	// 只获取读锁，多个goroutine可以同时读数据，不再相互等待。
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func main() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	for i := 0; i < 10; i++ {
		go fmt.Printf("Sum = %d\n", readSum())
	}
	time.Sleep(2 * time.Second)
}
