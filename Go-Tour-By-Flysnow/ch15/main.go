package main

import "fmt"

func main() {
	// Go语言程序所管理的虚拟内存空间会被分为两部分：堆内存和栈内存
	// 因为大部分程序数据分配在堆内存上，所以一个程序的大部分内存占用也是在堆内存上。
	// 我们常说的Go语言的内存垃圾回收是针对堆内存的垃圾回收。

	// 如果在声明一个变量的时候就给这个变量赋值，这种操作就称为变量的初始化。
	// 如果要对一个变量赋值，这个变量必须有对应的分配好的内存，这样才可以对这块内存进行操作，达到赋值的目的。
	// 所以一个变量必须要经过声明、内存分配才能赋值，才可以进行初始化。
	var s string
	fmt.Printf("%#v, %p\n", s, &s)
	fmt.Println()

	// 指针类型在声明的时候，Go语言并没有自动分配内存，所以不能对其进行赋值操作，这与值类型不一样。
	var sp *string
	fmt.Printf("%#v, %p\n", sp, sp)
	//*sp = "abc" 会报错，读取 *sp 也会报错
	fmt.Println()

	// new 函数作用就是根据传入的类型申请一块内存，然后返回指向这块内存的指针，指针指向的数据就是该类型的零值。
	// new函数只用于分配内存，并且把内存清零，也就是返回一个指向对应类型零值的指针。new函数一般用于需要显式地返回指针的情况，不是太常用。
	sp = new(string)
	fmt.Printf("%#v, %p\n", *sp, sp)
	fmt.Println()

	// make函数只用于slice、chan和map这三种内置类型的创建与初始化。
	// 因为这三种类型的结构比较复杂，比如slice要提前初始化内部元素的类型、slice的长度和容量等，这样才可以更好地使用它们。
	c := make(chan string)
	fmt.Printf("%#v, %p\n", c, c)
}
