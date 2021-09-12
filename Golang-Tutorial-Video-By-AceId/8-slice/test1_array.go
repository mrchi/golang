// 固定长度数组
package main

import "fmt"

func main() {
	// 声明固定长度的数组
	// 不显式赋值时，使用类型默认值
	var myArray1 [3]int

	// 声明固定长度数组，赋值前 4 个元素
	myArray2 := [5]int{1, 2, 3, 4}

	fmt.Printf("myArray1: %v, type: %T\n", myArray1, myArray1)
	fmt.Printf("myArray2: %v, type: %T\n", myArray2, myArray2)

	// 数组遍历，类似 Python 中的 for i in enumerate(lst)
	for index, val := range myArray2 {
		fmt.Printf("index = %d, val = %v\n", index, val)
	}
	// 数组遍历，类似 Python 中的 for i in range(len(i))
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	// 固定长度数组作为参数传递时，是值拷贝，在函数中改变元素值不会影响原数组
	printArray(myArray2)
	fmt.Printf("myArray2: %v, type: %T\n", myArray2, myArray2)
}

func printArray(myArray [5]int) {
	myArray[4] = 100
	fmt.Printf("myArray in printArray(): %v, type: %T\n", myArray, myArray)
}
