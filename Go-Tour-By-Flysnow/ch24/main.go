package main

import (
	"fmt"
	"net"

	"golang.org/x/exp/constraints"
)

// 在Go语言中是通过[]来定义类型T的
// 标识T就是一个类型参数，在定义函数的时候，我们不知道它是什么类型，但是在调用函数的时候，会被替换为具体的参数类型，这也是它被称为类型参数的原因
func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// 类型参数不仅可以用于函数，也可以用于类型中，当用于类型中的时候，我们就得到了一个泛型类型。
type GSlice[T constraints.Ordered] []T

// 泛型类型可以有方法，并且方法接收者的类型参数要与泛型类型的数量一样
// 唯一不同的是可以省略类型的约束，因为已经在泛型类型声明的时候，通过类型参数约束过了
// 尽管方法可以使用泛型类型的参数，但是方法本身并不能提供额外的类型参数，所以如果你需要额外的类型参数的话，只能通过函数实现。
func (s GSlice[T]) filter(f func(T) bool) GSlice[T] {
	result := GSlice[T]{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// 泛型类型是可以自引用的，如果泛型类型有多个类型参数，自引用要保持顺序一致。
type P[T1, T2 any] struct {
	F *P[T1, T2]
}

// 类型约束本身就是一个interface，所以可以用 interface 进行约束
func stringify[T fmt.Stringer](s []T) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[i] = v.String()
	}
	return result
}

// 用作约束的接口可以被赋予名称（例如Ordered），或者它们可以是内联在类型参数列表中的接口
// 作为类型集的接口是一种强大的新机制，是使类型约束在Go中起作用的关键。
func test1[S interface{ ~[]E }, E interface{}](s S) {
	fmt.Printf("%#v\n", s)
}

func test2[S ~[]E, E any](s S) {
	fmt.Printf("%#v\n", s)
}

type Point []int

func (p Point) String() string {
	return fmt.Sprintf("Point: %#v", p)
}

func scale[S ~[]E, E constraints.Integer](s S) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * 2
	}
	return r
}

func main() {
	// -------------------- 类型参数 --------------------
	// 泛型函数实例化：向函数提供类型参数的这种方式称为泛型函数实例化
	// 1)编译器把泛型函数中的类型参数替换为真正的类型实参
	// 2)然后验证这个真正的类型实参是否满足函数定义的类型约束
	print[int]([]int{1, 2, 3})
	fmt.Println()

	// 一个泛型函数一旦被实例化后，它就是一个非泛型函数了，它的使用就与普通的函数一样
	p := print[string]
	p([]string{"a", "b", "c"})
	fmt.Println()

	// -------------------- 泛型类型 --------------------
	s := GSlice[int]{1, 2, 3}
	fmt.Printf("%#v\n", s.filter(func(i int) bool { return i <= 2 }))
	fmt.Println()

	// -------------------- 类型约束 --------------------
	fmt.Printf("%#v\n", stringify([]net.IP{net.IPv4(192, 168, 1, 1), net.IPv4(255, 255, 255, 255)}))
	fmt.Println()

	// any可以表示任意类型，它是一个空接口，也就是interface{}，即 any 是 interface{} 的类型别名。
	// 对于泛型，可以把约束视为类型参数的元类型

	// -------------------- 类型约束·类型集 --------------------
	// 类型参数列表的每个类型参数都有一个类型。因为类型参数本身就是一种类型，所以类型参数的类型定义了类型集。这种元类型称为类型约束。
	// 在Go中，类型约束必须是接口。也就是说，接口类型既可以作为值类型，也可以作为元类型

	// 对于接口的定义，另一种看法是，接口定义了一组类型，即实现这些方法的类型。从这个角度来看，作为接口类型集元素的任何类型都实现了该接口
	// 在Go1.18版本中，接口除了可以像以前一样包含方法和嵌入式接口外，也可以嵌入非接口类型、联合和底层类型集。

	test1([]int{1, 2, 3})
	test2([]string{"1", "2", "3"})
	fmt.Println()

	// -------------------- 类型推导 --------------------
	// 从函数参数的类型推导出类型参数的情形称为函数参数类型推导。
	// 函数实参类型推导仅适用于函数参数中使用的类型参数，不适用于仅用于函数结果或仅在函数体中的类型参数。
	// 例如，它不适用于像MakeT[T any]()T这样只使用T作为结果的函数。

	// 约束类型推导。约束类型推导从类型形参约束中推导出类型实参，通常情况是：当一个约束对某种类型使用～type形式时，该类型是使用其他类型形参编写的。
	// 如果类型推导失败，编译器将给出错误消息，在这种情况下，我们可以只提供必要的类型参数。
	point1 := Point{1, 2, 3}
	fmt.Println(scale(point1).String())
}
