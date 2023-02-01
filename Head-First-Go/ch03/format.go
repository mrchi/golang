package main

import "fmt"

func main() {
	// 格式化动词
	fmt.Printf("%d, %f, %t\n", 3, 3.134, false)
	fmt.Printf("%v, %#v\n", "\t", "\t") // %# 相当于 Python 中的 repr
	fmt.Printf("%v, %T\n", "hello", "hello")
	fmt.Printf("%%\n")

	// 格式化宽度，%[最小宽度].[小数宽度][格式]
	fmt.Println("----------")
	fmt.Printf("%10s\n", "abc")
	fmt.Printf("%10d\n", 123)
	fmt.Printf("%10.5f\n", 3.14)
	fmt.Printf("%.2f\n", 3.1415926)
}
