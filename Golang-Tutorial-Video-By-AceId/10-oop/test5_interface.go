// interface{}
package main

import "fmt"

// interface{} 可以作为万能数据类型
func PrintInterface(arg interface{}) {
	// 数据类型断言，类似 Python 中的 isinstance
	value, ok := arg.(Game)
	fmt.Printf("arg = %v, type = %T\n", arg, arg)
	fmt.Printf("value = %v, is Game: %v\n", value, ok)
}

type Game struct {
	author string
}

func main() {
	book := Game{"Halo"}

	PrintInterface(book)
	PrintInterface("abc")
	PrintInterface(123)
	PrintInterface(3.1415926)
}
