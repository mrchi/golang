// slice 的 4 种声明方式
package main

import "fmt"

func main() {
	// 声明 slice，并且初始化默认值为 1, 2, 3，长度为 3
	slice1 := []int{1, 2, 3}
	fmt.Printf("slice1 is %v\n", slice1)

	// 声明空 slice，空 slice == nil 返回 true
	var slice2 []int
	fmt.Printf("slice2 is %v\n", slice2)
	fmt.Printf("slice2 == nil: %t\n", slice2 == nil)

	// 开辟 3 个长度的空间
	slice2 = make([]int, 3)
	fmt.Printf("slice2 is %v\n", slice2)

	// 开辟空间后才能对元素赋值
	slice2[2] = 10
	fmt.Printf("slice2 is %v\n", slice2)

	// 声明 slice，分配 4 个长度的空间，元素值为类型默认值
	var slice3 []int = make([]int, 4)
	fmt.Printf("slice3 is %v\n", slice3)

	// 声明 slice，分配 4 个长度的空间，元素值为类型默认值
	slice4 := make([]int, 4)
	fmt.Printf("slice4 is %v\n", slice4)
}
