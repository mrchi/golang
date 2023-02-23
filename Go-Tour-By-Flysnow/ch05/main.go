package main

import "fmt"

// 形参类型一致时，我们可以省略其中一个类型的声明
func sum1(a, b int) int {
	return a + b
}

// 可以为返回值命名
// 直接为命名的返回值赋值，也就等于函数有了返回值，所以可以 return 空
func sum2(a, b int) (result int) {
	result = a + b
	return
}

// 但如果 return 返回了其他值，return 的优先级更高
func sum3(a, b int) (result int) {
	result = a + b
	return 123
}

// 定义可变参数，只要在参数类型前加三个点
func sum4(init int, a ...int) (result int) {
	result = init
	// 可变参数的类型其实就是切片
	for _, v := range a {
		result += v
	}
	return
}

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// ----- 方法 -----

type Age uint

// 定义方法不要混用指针类型接收者和变量调用值类型接收者，此处只是示例。
func (a Age) String() {
	fmt.Printf("The age is %d\n", a)
}

func (a *Age) Modify(newAge uint) {
	*a = Age(newAge)
}

func main() {
	// 函数参数和返回值
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(1, 2))
	fmt.Println(sum4(100, 1, 2, 3, 4))
	fmt.Println()

	// 匿名函数和闭包
	cl := closure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println()

	// 在调用方法的时候，传递的接收者本质上都是副本，只不过一个是这个值的副本，一个是指向这个值的指针的副本
	var a Age
	p := &a

	// 如果使用一个值类型变量调用指针类型接收者的方法，Go语言编译器会自动帮我们取指针调用，以满足指针接收者的要求。
	a.Modify(3)
	a.String()

	p.Modify(2)
	// 同样的原理，如果使用一个指针类型变量调用值类型接收者的方法，Go语言编译器会自动帮我们解引用调用，以满足值类型接收者的要求。
	p.String()
}
