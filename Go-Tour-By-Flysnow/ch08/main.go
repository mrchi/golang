package main

import (
	"fmt"
	"time"
)

// 单向管道的声明，只需要在声明的时候带上＜-操作符即可。
// 在函数或者方法的参数中，使用单向管道的较多，这样可以防止一些操作对管道的影响
func downloadFile(out chan<- int, timeout int) {
	time.Sleep(time.Duration(timeout) * time.Second)
	out <- timeout
}

func main() {
	// channel内部使用了互斥锁来保证并发的安全

	// 定义有缓冲管道，容量为 3
	ch := make(chan int, 3)
	// 通过内置函数cap可以获取管道的容量，通过内置函数len可以获取管道中元素的个数。
	fmt.Printf("Length=%d, Capacity=%d\n", len(ch), cap(ch))
	fmt.Println()

	// 关闭管道，关闭后就不能向里面发送数据了
	ch <- 1
	close(ch)

	// 但是我们还可以接收管道里的数据，如果管道里没有数据的话，接收的数据是元素类型的零值。
	fmt.Printf("Value from closed channel: %#v\n", <-ch)
	fmt.Printf("Value from closed channel: %#v\n", <-ch)
	fmt.Println()

	// select + channel
	// 在Go语言中，通过select语句可以实现多路复用
	// 多路复用可以简单地理解为，在N个channel中，任意一个channel有数据产生，select都可以监听到，然后执行相应的分支，接收数据并处理。
	// 如果同时有多个case可以被执行，则随机选择一个
	// 如果一个select没有任何case可以被执行，那么它会一直等待下去。
	firstCh := make(chan int)
	secondCh := make(chan int)
	thirdCh := make(chan int)

	go downloadFile(firstCh, 1)
	go downloadFile(secondCh, 2)
	go downloadFile(thirdCh, 3)

	for i := 0; i < 3; i++ {
		select {
		case filePath := <-firstCh:
			fmt.Println(filePath)
		case filePath := <-secondCh:
			fmt.Println(filePath)
		case filePath := <-thirdCh:
			fmt.Println(filePath)
		}
	}
}
