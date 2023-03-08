package main

import "fmt"

const name = "张三"

// 指针作为函数返回值的时候，一定会发生逃逸。
func NewString() *string {
	s := new(string)
	*s = "李四"
	return s
}

func NewString2() string {
	s := new(string)
	*s = "李四"
	return *s
}

func main() {
	NewString()
	NewString2()

	// 被已经逃逸的指针引用的变量也会发生逃逸。
	fmt.Println("Hello")

	// Go语言中有3个比较特殊的类型，它们是slice、map和chan，被这三种类型引用的指针也会发生逃逸。
	m := map[int]*string{}
	s := "张三"
	m[0] = &s

	// 逃逸分析是判断变量是分配在堆内存上还是栈内存上的一种方法，在实际的项目中要尽可能避免逃逸，这样就不会被GC拖慢速度，从而提升效率。

	// 总结几个优化的小技巧
	// 1)第1个需要介绍的技巧是尽可能避免逃逸，因为栈内存效率更高，还不用GC。
	// 2)如果避免不了逃逸，还是在堆上分配了内存，那么对于频繁的内存申请操作，我们要学会重用内存，比如使用sync.Pool，这是第2个技巧。
	// 3)第3个技巧就是选用合适的算法，达到高性能的目的，比如以空间换时间。
	// 除此之外，还有一些小技巧，比如要尽可能避免使用锁、并发加锁的范围要尽可能小、使用StringBuilder做string和[]byte之间的转换、defer嵌套不要太多等。

	// Go语言自带的性能剖析工具pprof
}
