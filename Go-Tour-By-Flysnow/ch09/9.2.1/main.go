package main

import (
	"fmt"
	"sync"
	"time"
)

var sum = 0
var mutex sync.Mutex // 互斥锁mutex

func add(i int) {
	// Mutex的Lock和Unlock方法总是成对出现
	mutex.Lock()
	// 而且要确保执行Lock获得锁后，一定执行UnLock释放锁，所以在函数或者方法中会采用defer语句释放锁
	defer mutex.Unlock()
	// 临界区指的是一个访问共享资源的程序片段，而这些共享资源又有无法同时被多个协程访问的特性。
	// 被加锁保护的代码片段称为临界区
	sum += i
}

func main() {
	for i := 0; i < 100; i++ {
		go add(10)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Sum = %d\n", sum)
}
