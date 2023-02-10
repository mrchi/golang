package main

import (
	"custom"
	"fmt"
	"gadget"
)

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
	fmt.Println()

	// error interface, 需要实现 Error() 方法
	err := custom.CheckTemperature(20.1, 20.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// fmt.Stringer interface, 需要实现 String() 方法
	nestle := custom.CoffeePot("Nestle")
	fmt.Println(nestle)
	fmt.Print(nestle, "\n")
	fmt.Printf("%v\n", nestle)
	fmt.Println()

	// Interface{}类型称为空接口，用来接收任何类型的值。
	// 不需要实现任何方法来满足空接口，所以所有的类型都满足它。
	custom.AcceptAnything(gadget.TapePlayer{})
	custom.AcceptAnything(gadget.TapeRecorder{})
	custom.AcceptAnything("Hello")
}
