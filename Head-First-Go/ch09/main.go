package main

import "fmt"

type FloatA float64
type FloatB float64

func main() {
	// 具有相同基础类型的类型之间可以相互转换
	float_b := FloatB(FloatA(33.3))
	fmt.Printf("%#v, %T\n", float_b, float_b)
	fmt.Println()

	// 即使具有相同的基础类型，也认为是不同类型，不能跨类型赋值
	// 这会报错 float_b = FloatA(33.3)

	// 定义类型提供与基础类型相同的运算
	fmt.Println(FloatA(3.0) < FloatA(3.1))
	fmt.Println(FloatA(3.0) > FloatA(3.1))
	fmt.Println(FloatA(3.0) == FloatA(3.1))
	fmt.Println(FloatA(3.0) + FloatA(3.0))
	fmt.Println(FloatA(3.0) - FloatA(3.0))
	fmt.Println(FloatA(3.0) * FloatA(3.0))
	fmt.Println(FloatA(3.0) / FloatA(3.0))
	fmt.Println()

	// 定义类型可以被用来与字面值一起用于运算
	fmt.Println(FloatA(3.0) + 3.0)
	fmt.Println()

	// 定义类型不能用来与不同类型的值一起运算，即使它们是来自相同的基础类型。
	// 这会报错 FloatA(3.0) > FloatB(3.0)

}
