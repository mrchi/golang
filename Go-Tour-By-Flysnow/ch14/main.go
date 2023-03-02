package main

import "fmt"

type address struct {
	province string
	city     string
}

func (a address) String() string {
	return fmt.Sprintf("The addr is %s%s", a.province, a.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	// 当值类型作为接收者实现了某接口时，它的指针类型也同样实现了该接口。
	addr := address{province: "北京", city: "北京"}
	fmt.Printf("%#v, %T\n", addr, addr)
	fmt.Printf("%#v, %T\n", &addr, &addr)
	printString(addr)
	printString(&addr)
	fmt.Println()

	// 虽然指向具体类型的指针可以实现一个接口，但是指向接口的指针永远不可能实现该接口。
	var si fmt.Stringer = address{province: "上海", city: "上海"}
	fmt.Printf("%#v, %T\n", si, si)
	fmt.Printf("%#v, %T\n", &si, &si)
	printString(si)
	// printString(&si) 会报错
	fmt.Println()

	// Go语言中的函数传参都是值传递

	// 在Go语言中，任何创建map的代码（不管是字面量还是make函数）最终调用的都是runtime.makemap函数。
	// makemap函数返回的是一个*hmap类型，也就是说返回的是一个指针，所以我们创建的map其实就是*hmap类型。
	// Go语言通过make函数或字面量的包装为我们省去了指针的操作，让我们可以更容易地使用map。其实这就是语法糖
	m := make(map[string]int)
	m["test"] = 1
	fmt.Printf("%#v %p\n", m, m)
	// 在参数传递时，它还是值传递，并不是其他编程语言中所谓的引用传递。
	func(m map[string]int) {
		m["test"] = 2
	}(m)
	fmt.Printf("%#v %p\n", m, m)
	fmt.Println()

	// 创建的chan其实是*hchan，也可以称为引用类型
	c := make(chan string, 5)
	fmt.Printf("%#v %p\n", c, c)
	fmt.Println()

	// 可以称为引用类型的变量的零值都是nil，包括 slice/map/指针/函数/channel/interface
	var i interface{}
	fmt.Printf("%#v %p\n", i, i)
	fmt.Println()
}
