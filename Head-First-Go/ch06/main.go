package main

import "fmt"

func main() {
	// 只声明，长度为 0
	var mySlice []int
	fmt.Printf("%v, %#v, length %d\n", mySlice, mySlice, len(mySlice))

	// make 函数赋值
	mySlice = make([]int, 3)
	mySlice[0] = 1
	mySlice[1] = 2
	fmt.Printf("%v, %#v, length %d\n", mySlice, mySlice, len(mySlice))

	// 短变量声明
	mySlice2 := make([]int, 3)
	fmt.Printf("%v, %#v, length %d\n", mySlice2, mySlice2, len(mySlice2))

	// 切片字面量
	mySlice3 := []int{1, 2, 3}
	fmt.Printf("%v, %#v, length %d\n", mySlice3, mySlice3, len(mySlice3))

	// 从数组创建索引
	myArray := [5]int{0, 1, 2, 3, 4}
	mySlice4 := myArray[1:4]
	mySlice5 := myArray[1:]
	mySlice6 := myArray[:4]
	fmt.Printf("%#v\n", mySlice4)
	fmt.Printf("%#v\n", mySlice5)
	fmt.Printf("%#v\n", mySlice6)

	// 切片是数组的视图
	myArray[3] = 99
	fmt.Println(myArray, mySlice4, mySlice5, mySlice6)
	mySlice4[2] = 88
	fmt.Println(myArray, mySlice4, mySlice5, mySlice6)

	// append 后的底层数组可能与原数组不同
	myArray2 := [3]string{"a", "b", "c"}
	mySlice7 := myArray2[:2]
	fmt.Printf("Array: %#v, Slice: %#v\n", myArray2, mySlice7)
	mySlice7 = append(mySlice7, "X", "Y")
	fmt.Printf("Array: %#v, Slice: %#v\n", myArray2, mySlice7)

	fmt.Printf("%#v\n", inRange(2, 4, 1, 2, 3, 4, 5))
	mySlice8 := []float64{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", inRange(2, 4, mySlice8...))
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}
