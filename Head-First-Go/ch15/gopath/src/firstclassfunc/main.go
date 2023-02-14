// Go语言支持一级函数
package main

import "fmt"

func SayHi() {
	fmt.Println("Hi")
}

// 接受函数作为参数的函数还需要指定传入函数应该具有的参数和返回类型
func Divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func RunTwice(function func()) {
	function()
	function()
}

func main() {
	// 在具有一级函数的编程语言中，可以将函数分配给变量，然后从这些变量调用函数。
	greeterFunction := SayHi
	// 函数的参数和返回值是其类型的一部分。
	// 保存函数的变量需要指定函数应该具有哪些参数和返回值。
	// 该变量只能保存参数的数量和类型以及返回值与指定类型匹配的函数。
	var mathFunction func(int, int) float64

	// 具有一级函数的编程语言还允许将函数作为参数传递给其他函数。
	RunTwice(greeterFunction)
	mathFunction = Divide
	fmt.Println(mathFunction(3, 2))
}
