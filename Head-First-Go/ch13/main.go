package main

import (
	"fmt"
	"time"
)

func PrintA() {
	for i := 0; i < 50; i++ {
		fmt.Print("A")
	}
}

func PrintB() {
	for i := 0; i < 50; i++ {
		fmt.Print("B")
	}
}

// 每个Go程序的main函数都是使用goroutine启动的，因此每个Go程序至少运行一个goroutine。
func main() {
	// 要启动另一个goroutine，可以使用go语句
	// 在正常情况下，Go不能保证何时在goroutine之间切换，或者切换多长时间
	// 我们不能在go语句中使用函数返回值
	go PrintA()
	go PrintB()
	time.Sleep(3 * time.Second)
	fmt.Println()
	fmt.Println("End main()")
}
