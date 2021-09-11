// defer
package main

import "fmt"

func justReturn() int {
	fmt.Println("justReturn running...")
	return 0
}

func returnAndDefer() int {
	// defer 语句会在函数结束前，return 之后执行
	// 可以有多个 defer 语句，后定义的先执行（LIFO）
	defer fmt.Println("defer1 running...")
	defer fmt.Println("defer2 running...")

	return justReturn()
}

func main() {
	returnAndDefer()
}
