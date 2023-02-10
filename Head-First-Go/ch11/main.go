package main

import "fmt"

type MemUsage int

func (n *MemUsage) Double() {
	*n *= 2
}

type NetUsage int

func (n *NetUsage) Double() {
	*n *= 2
}

type Number interface {
	Double()
}

func main() {
	var number Number

	// 如果一个类型声明了指针接收器方法，你就只能将那个类型的指针传递给接口变量。
	n1 := MemUsage(1)
	number = &n1
	number.Double()
	fmt.Printf("%#v, %#v\n", number, n1)
}
