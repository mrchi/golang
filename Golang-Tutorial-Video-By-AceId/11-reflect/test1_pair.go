// 变量的 pair 结构
package main

import "fmt"

func main() {
	// 每个变量都是一个 type 和 value 组成的 pair
	a := "calm"

	var allType interface{}
	// pair 在被赋值时会被持续的传递
	allType = a
	fmt.Printf("allType = %v, type = %T\n", allType, allType)

	// 但是还拥有 interface 的特性，能进行断言
	value, ok := allType.(string)
	fmt.Printf("allType is string: %v, value: %v\n", ok, value)
}
