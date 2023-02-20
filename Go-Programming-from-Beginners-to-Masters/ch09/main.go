package main

import "fmt"

type myInt int
type myFloat float64
type myString string

const (
	a = 5
	b = 3.14
	c = "Hello"
)

func main() {
	// 字面值无须显式类型转换就可以直接赋值给对应的自定义类型的变量
	var j myInt = 5
	var f myFloat = 3.14
	var s myString = "Hello"
	fmt.Printf("%T %#v\n", j, j)
	fmt.Printf("%T %#v\n", f, f)
	fmt.Printf("%T %#v\n", s, s)
	fmt.Println()

	// 无类型常量在参与变量赋值和计算过程时无须显式类型转换
	var x myInt = a
	var y myFloat = b
	var z myString = c
	fmt.Printf("%T %#v\n", x, x)
	fmt.Printf("%T %#v\n", y, y)
	fmt.Printf("%T %#v\n", z, z)
	fmt.Println()

	// 无类型常量也拥有自己的默认类型
	// 无类型的布尔型常量、整数常量、字符常量、浮点数常量、复数常量、字符串常量对应的默认类型分别为bool、int、int32(rune)、float64、complex128和string。
	u := a                // 赋值给变量
	var i interface{} = a // 赋值给接口类型

	fmt.Printf("%T %#v\n", u, u)
	fmt.Printf("%T %#v\n", i, i)
	fmt.Println()

	i = c
	fmt.Printf("%T %#v\n", i, i)
	fmt.Println()

}
