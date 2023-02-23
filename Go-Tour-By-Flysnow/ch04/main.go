package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 长度也是数组类型的一部分，[5]string和[4]string不是同一种类型
	var a [4]string
	var b [5]string
	fmt.Printf("%T %T\n", a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
	fmt.Println()

	// 数组只针对特定索引元素初始化
	c := [5]string{1: "a", 3: "b", 4: "c"}
	fmt.Printf("%#v\n", c)
	fmt.Println()

	// 切片是一个具备三个字段的数据结构，分别是指向数组的指针data、长度len和容量cap。
	// 切片的容量不能比切片的长度小。
	// 容量是切片的空间，当切片的长度要超过容量的时候会进行扩容
	slice1 := make([]string, 3, 5)
	fmt.Printf("%#v, %d, %d", slice1, len(slice1), cap(slice1))
	fmt.Println()

	// 通过字面量初始化的切片，其长度和容量相同。
	slice2 := []string{"a", "b", "c"}
	fmt.Printf("%#v, %d, %d\n", slice2, len(slice2), cap(slice2))
	fmt.Println()

	// map 取值可以返回 1 个或 2 个值，第二个值标记该Key是否存在，如果存在，它的值为true
	m := map[string]int{"China": 123}
	fmt.Printf("%#v, %#v\n", m["China"], m["USA"])
	v, ok := m["China"]
	fmt.Printf("China: %#v, Exists: %#v\n", v, ok)
	v, ok = m["USA"]
	fmt.Printf("USA: %#v, Exists: %#v\n", v, ok)
	fmt.Println()

	// 字符串string也是一个不可变的字节序列，所以可以直接转为字节切片[]byte
	s := "abcd中国"
	fmt.Println([]byte(s))
}
