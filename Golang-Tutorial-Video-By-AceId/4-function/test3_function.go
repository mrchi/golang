// 函数定义和多返回值
package main

import "fmt"

// 返回单个返回值
func foo1(a string, b int) int {
	fmt.Printf("a = %s, b = %d\n", a, b)

	c := 100
	return c
}

// 返回多个返回值
func foo2() (int, string) {
	return 666, "abcdef"
}

// 返回多个返回值，有形参名称
func foo3() (r1 int, r2 string) {
	// 给返回值变量赋值，然后直接 return 即可
	r1 = 100
	r2 = "abc"
	return
}

// 返回多个返回值，形参是相同类型时的简写
func foo4() (r1, r2 int) {
	r1, r2 = 100, 200
	return
}

// 返回多个返回值，但没有给形参赋值时，形参使用类型的默认值
func foo5() (r1 int, r2 bool) {
	return
}

func main() {
	c := foo1("abc", 123)
	fmt.Printf("foo1 return: %d\n", c)

	ret1, ret2 := foo2()
	fmt.Printf("foo2 return: %d, %s\n", ret1, ret2)

	ret3, ret4 := foo3()
	fmt.Printf("foo3 return: %d, %s\n", ret3, ret4)

	ret5, ret6 := foo4()
	fmt.Printf("foo4 return: %d, %d\n", ret5, ret6)

	ret7, ret8 := foo5()
	fmt.Printf("foo5 return: %d, %t\n", ret7, ret8)
}
