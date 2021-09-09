// 变量声明方式
package main

import "fmt"

var g int = 10

// 全局变量声明不能省略 var 关键字
// g := 10

func main() {
	// 仅声明
	var a int
	fmt.Printf("a = %d\n", a)

	// 声明并赋值
	var b int = 10
	fmt.Printf("b = %d\n", b)
	fmt.Printf("Type of b = %T\n", b)

	// 自动推断类型
	var c = "abcd"
	fmt.Printf("c = %s\n", c)

	// 省略声明 var 关键字
	d := 100
	fmt.Printf("d = %d\n", d)
	fmt.Printf("Type of d = %T\n", d)

	// 使用全局变量
	fmt.Printf("g = %d\n", g)

	// 多变量声明
	var aa, bb int = 10, 20
	fmt.Printf("aa = %d, bb = %d\n", aa, bb)
	var cc, dd = 100, "abcd"
	fmt.Printf("cc = %d, dd = %s\n", cc, dd)

	// 多行多变量声明
	var (
		ee int  = 10
		ff bool = true
	)
	fmt.Printf("ee = %d, ff = %t\n", ee, ff)
}
