package main

import (
	"fmt"
	"sync"
)

var sum = 0
var mutex sync.RWMutex

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func main() {
	// 声明一个sync.WaitGroup
	var wg sync.WaitGroup

	// 通过Add方法设置计数器的值，需要跟踪多少个协程就设置多少
	wg.Add(110)

	for i := 0; i < 100; i++ {
		go func() {
			// 在每个协程执行完毕后调用Done方法，让计数器减1
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			// 在每个协程执行完毕后调用Done方法，让计数器减1
			defer wg.Done()
			fmt.Printf("Sum = %d\n", readSum())
		}()
	}

	// 最后调用Wait方法一直等待，直到计数器值为0，也就是所有跟踪的协程都执行完毕。
	wg.Wait()

	fmt.Printf("All Done. Sum = %d\n", readSum())
}
