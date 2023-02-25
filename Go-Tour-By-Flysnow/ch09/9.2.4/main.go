package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Once适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景。
	var once sync.Once
	var wg sync.WaitGroup

	onceBody := func() {
		fmt.Println("Only once")
	}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 用了once.Do方法，所以函数onceBody只会被执行一次
			// 该函数f必须没有返回值没有参数：f func()
			once.Do(onceBody)
		}()
	}

	wg.Wait()
}
