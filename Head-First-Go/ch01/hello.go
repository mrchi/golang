package main

import (
	"fmt"
	"reflect"
)

func main() {
	// hello world
	fmt.Println("Hello, world!")

	// 字符串和符文（rune）
	fmt.Println("你", '你')

	// 变量声明和赋值
	var name string
	name = "gopher"
	fmt.Println("Hello", name)

	// 声明变量并赋值
	var quantity int = 4
	var length, width = 3.14, 3.15
	fmt.Println("Quantity", quantity, "\t", "Area", length*width)

	// 短变量声明
	customerName := "Liu Peiqiang"
	fmt.Println(customerName, reflect.TypeOf(customerName))

	// 零值
	var age int
	fmt.Println(reflect.TypeOf(age), age)

	// 类型转换
	var price int = 30
	var tax_rate float64 = 0.02
	fmt.Println("Tax is", float64(price)*tax_rate)
}
