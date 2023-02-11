package main

import "fmt"

func one() {
	defer fmt.Println("Defer in one()")
	two()
}

func two() {
	defer fmt.Println("Defer in two()")
	three()
}

func three() {
	defer fmt.Println("Defer in three()")
	// 当程序发生panic时，panic输出中包含堆栈跟踪，即调用堆栈列表。
	panic("This is a panic")
}

// 当程序出现panic时，所有延迟的函数调用仍然会被执行。
// 如果有多个延迟调用，它们的执行顺序将与被延迟的顺序相反。
func main() {
	one()
}
