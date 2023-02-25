package main

import (
	"fmt"
	"time"
)

var sum = 0

func add(i int) {
	sum += i
}

func main() {
	// 所有 goroutine 中都在访问 sum，存在资源竞争问题，不是并发安全的
	// 使用go build、go run、go test这些Go语言工具链提供的命令时，添加-race标识可以帮你检查Go语言代码是否存在资源竞争
	for i := 0; i < 100; i++ {
		go add(10)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Sum = %d\n", sum)
}
