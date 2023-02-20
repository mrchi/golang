package main

import "fmt"

// 隐式重复前一个非空表达式
const (
	Apple, Banana = 11, 22
	Strawberry, Grape
	Pear, Watermelon
)

// iota是Go语言的一个预定义标识符，它表示的是const声明块（包括单行声明）中每个常量所处位置在块中的偏移值（从零开始）
// 同时，每一行中的iota自身也是一个无类型常量
const (
	Flag1       = 1 << iota // iota = 0，因此 Flag1 = 1
	Flag2                   // iota = 1，同时隐式重复前一个非空表达式，Flag2 = 1 << 1 = 2
	Flag3                   // iota = 2，同时隐式重复前一个非空表达式，Flag2 = 1 << 2 = 4
	AnotherFlag = iota      // iota = 3，因此 AnotherFlag = 3
	LastFlag                // iota = 4，同时隐式重复前一个非空表达式，因此 LastFlag = 4
)

// 位于同一行的iota即便出现多次，其值也是一样的
const (
	Orange, Cherry = iota, iota + 3.14 // Go的枚举常量不限于整型值，也可以定义浮点型的枚举常量
	_, _                               // 略过某一个 iota 值
	FruitX, FruitY
)

func main() {
	fmt.Printf("%#v, %#v, %#v, %#v, %#v, %#v\n", Apple, Banana, Strawberry, Grape, Pear, Watermelon)
	fmt.Println()

	fmt.Printf("%#v, %#v, %#v, %#v, %#v\n", Flag1, Flag2, Flag3, AnotherFlag, LastFlag)
	fmt.Println()

	fmt.Printf("%#v, %#v\n", Orange, Cherry)
	fmt.Printf("%#v, %#v\n", FruitX, FruitY)
	fmt.Println()

}
