package main

import (
	"fmt"
	"fuel"
)

type FloatA float64
type FloatB float64
type Number float64

// 【这只是例子】为了一致性，你所有的类型函数接受值类型，或者都接受指针类型，但是你应该避免混用的情况。
func (n *Number) Double() {
	*n *= 2
}

// 【这只是例子】为了一致性，你所有的类型函数接受值类型，或者都接受指针类型，但是你应该避免混用的情况。
func (n Number) Print() {
	fmt.Printf("%#v\n", n)
}

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

	// 调用定义的方法
	carFuel := fuel.Gallons(100)
	fmt.Printf("%#v\n", carFuel.ToLiters().ToGallons())
	fmt.Printf("%#v\n", carFuel.ToMilliliters().ToGallons())
	fmt.Println()

	// 指针类型作为方法接收器参数，用非指针类型调用，Go 自动转换
	number := Number(3.14)
	number.Double()
	fmt.Printf("%#v\n", number)
	// 非指针类型作为方法接收器参数，用指针类型调用，Go 自动转换
	pointer := &number
	pointer.Print()

}
