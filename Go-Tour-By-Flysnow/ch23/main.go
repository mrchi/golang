package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Go语言为我们提供了自定义类型约束的能力。
// 这里用的“|”符号表示“或”，也就是并集的意思。
// ～符号，它表示对派生出的类型也有用。
type Number interface {
	~int | float32 | float64 | int32
}

// 泛型函数定义：它在函数名和小括号形参之间，多了一个方括号，用于声明类型参数。
func min[V Number](a, b V) V {
	if a < b {
		return a
	} else {
		return b
	}
}

// 内置的类型约束，在 golang.org/x/exp/constraints 包中
// [Signed] 所有有符号整型的并集
// [Unsigned] 无符号型的约束
// [Integer] Signed和Unsigned约束，把它俩并起来就可以了。
// [Float] Float就是把 float32 和 float64 合并起来。
// [Ordered] 多了一个string，这是因为string也是有序的，即支持<、<=、>=、>操作符
func max[V constraints.Ordered](a, b V) V {
	if a < b {
		return b
	} else {
		return a
	}
}

type MyInt int

// 泛型类型的切片
// Go SDK还提供了两种约束类型：any和comparable，第一个是一个空接口，表示任意类型；第二个是可比较的接口，可用于map类型的key键约束。
type GSlice[T any] []T

// 函数式编程 map 方法，采用类型方法的方式，让它更符合函数编程的链式调用。
func (s GSlice[T]) map1(f func(T) T) GSlice[T] {
	result := make(GSlice[T], len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// 函数式编程 reduce 方法
func (s GSlice[T]) reduce(f func(previousV, currentV T) T) T {
	var result T
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

// 函数式编程 filter 方法
func (s GSlice[T]) filter(f func(T) bool) GSlice[T] {
	result := GSlice[T]{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// find 方法
func (s GSlice[T]) find(f func(T) bool) (T, bool) {
	var result T
	for _, v := range s {
		if f(v) {
			result = v
			return result, true
		}
	}
	return result, false
}

func main() {
	// 在调用时，传递了一个具体的实参类型，它们使用方括号([])传递
	// Go语言的编译器，将类型参数替换为该函数调用时具体指定的类型
	fmt.Printf("%#v %T\n", min[int](1, 2), min[int](1, 2))

	// 通过Go编译器，它可以根据函数调用时实参的类型推导出所需的类型参数
	// 在调用泛型函数时，省略了类型参数后，就变得很简洁了，与原来的非泛型函数一样的用法，没有多余额外的编码工作。
	fmt.Printf("%#v %T\n", min(1, 2), min(1, 2))
	fmt.Printf("%#v %T\n", min(3.14, 2.55), min(3.14, 2.55))
	// 对派生出的类型也有用
	fmt.Printf("%#v %T\n", min(MyInt(3), MyInt(2)), min(MyInt(3), MyInt(2)))
	fmt.Println()

	fmt.Printf("%#v %T\n", max(1, 2), max(1, 2))
	fmt.Printf("%#v %T\n", max("abc", "bcd"), max("abc", "bcd"))
	fmt.Println()

	// 自定义的函数式编程方法
	m := GSlice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := m.filter(func(v int) bool {
		return v%2 == 0
	}).map1(func(v int) int {
		return v / 2
	}).reduce(func(pre, v int) int {
		return pre + v
	})
	fmt.Println(result)
	fmt.Println()

	result, ok := m.find(func(v int) bool {
		return v >= 3
	})
	fmt.Println(result, ok)
	result, ok = m.find(func(v int) bool {
		return v < 0
	})
	fmt.Println(result, ok)
	fmt.Println()

}
