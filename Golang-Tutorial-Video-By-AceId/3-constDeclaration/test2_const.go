// 常量定义和 iota
package main

import "fmt"

const (
	// const 中每行 iota 都会累加 1，第一行的 iota 的默认值是 0
	// iota 只能在 const 中使用
	BEIJING   = iota // 0
	SHANGHAI         // 1
	GUANGZHOU        // 2
	SHENZHEN         // 3
)

const (
	XIANGANG = iota * 10 // 0
	AOMEN                // 10
	TAIWAN               // 20
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = 1, b = 2
	c, d                      // iota = 1, c = 2, d = 3

	e, f = iota * 2, iota * 3 // iota = 2, e = 4, f = 6
	g, h                      // iota = 3, g = 6, h = 9
)

func main() {
	const length int = 10
	fmt.Printf("length = %d\n", length)

	fmt.Println(BEIJING, SHANGHAI, GUANGZHOU, SHENZHEN)
	fmt.Println(XIANGANG, AOMEN, TAIWAN)
	fmt.Println(a, b, c, d, e, f, g, h)
}
