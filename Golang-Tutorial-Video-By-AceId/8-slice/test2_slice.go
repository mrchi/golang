// 动态数组 slice
package main

import "fmt"

func printArray(myArray []int) {
	myArray[0] = 100
	fmt.Printf("myArray in printArray(): %v, type: %T\n", myArray, myArray)
}

func main() {
	// 定义 slice 并赋值元素，数组的大小是传入元素的个数
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray: %v, type is %T\n", myArray, myArray)

	// slice 作为参数传递时，为引用传递，在函数中改变元素会影响原 slice
	printArray(myArray)
	fmt.Printf("myArray: %v, type: %T\n", myArray, myArray)
}
