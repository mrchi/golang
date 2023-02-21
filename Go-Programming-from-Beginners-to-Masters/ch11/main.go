package main

import "fmt"

func main() {
	var a int
	fmt.Printf("%#v\n", a)
	var b float64
	fmt.Printf("%#v\n", b)
	var c bool
	fmt.Printf("%#v\n", c)
	var d string
	fmt.Printf("%#v\n", d)
	var e byte
	fmt.Printf("%#v\n", e)
	// 指针零值 nil
	var f *int
	fmt.Printf("%#v\n", f)
	// slice 零值 nil
	var g []int
	fmt.Printf("%#v\n", g)
	// channel 零值 nil
	var h chan int
	fmt.Printf("%#v\n", h)
	// map 零值 nil
	var i map[int]int
	fmt.Printf("%#v\n", i)
	fmt.Println()

	// Go的零值初始是递归的，即数组、结构体等类型的零值初始化就是对其组成元素逐一进行零值初始化。
	type Webhook struct {
		Url string
	}
	var w Webhook
	fmt.Printf("%#v\n", w)
	fmt.Println()

	// Go中的切片类型具备零值可用的特性，我们可以直接对其进行append操作，而不会出现引用nil的错误。
	// 不过，Go并非所有类型都是零值可用的，并且零值可用也有一定的限制
	// 像map这样的原生类型也没有提供对零值可用的支持
	// 另外零值可用的类型要注意尽量避免值复制。
	var zeroSlice []int
	zeroSlice = append(zeroSlice, 1)
	fmt.Printf("%#v\n", zeroSlice)
	fmt.Println()
}
