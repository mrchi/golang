package main

import "fmt"

func main() {
	// 当switch之后有表达式时，case后的值就要与这个表达式的结果类型相同
	switch i := 1; i {
	case 1:
		fmt.Println("1")
		fallthrough // fallthrough 到紧跟的下一个 case，即使不匹配也会执行，因此这里输出 1 和 2
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}
