package main

import "fmt"

func main() {
	var array1 [3]int
	array1[0] = 1
	fmt.Println(array1)
	fmt.Printf("%#v\n", array1)

	// 数组字面量
	var array2 [3]float64 = [3]float64{1, 2, 3}
	fmt.Printf("%#v\n", array2)
	var array3 = [3]float64{1, 2, 3}
	fmt.Printf("%#v\n", array3)
	array4 := [3]float64{1, 2, 3}
	fmt.Printf("%#v\n", array4)

	var v1 [3]int
	v2, v3 := 2, 3
	v1[2] = v2
	fmt.Println(v1, v2, v3)
}
