package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func printSliceInfo(s []string) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("[printSliceInfo]%#v\n", sh)
}

func printArrayInfo(a [2]string) {
	fmt.Printf("[printArrayInfo]%p\n", &a)
}

func main() {
	// append自动扩容的原理是新创建一个底层数组，把原来切片内的元素拷贝到新数组中，然后再返回一个指向新数组的切片。
	ss := []string{"张三", "李四"}
	fmt.Println(len(ss), cap(ss))
	ss = append(ss, "王五", "马六")
	fmt.Println(len(ss), cap(ss))
	fmt.Println()

	// SliceHeader是切片在运行时的表现形式，它有三个字段Data、Len和Cap。
	// 1)Data用来指向存储切片元素的数组。
	// 2)Len代表切片的长度。
	// 3)Cap代表切片的容量。
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("%#v, %T\n", sh, sh)
	fmt.Println()

	// 以此来证明不同切片对应的底层Data指向的可能是同一个数组。
	a1 := [2]string{"张三", "李四"}
	s1 := a1[0:1]
	s2 := a1[:]
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Println()

	// 切片的本质是SliceHeader，又因为函数的参数是值传递，所以传递的是SliceHeader的副本，而不是底层数组的副本
	printSliceInfo(s2)
	fmt.Println()

	// SliceHeader的三个字段的类型分别是uintptr、int和int，在64位的机器上，这三个字段最多也就是int64类型，一个int64占8字节，三个int64占24字节内存。
	fmt.Println(unsafe.Sizeof(s2))
	fmt.Println()

	// 和 map 相比，数组和切片的取值和赋值操作要更高效，因为它们是连续的内存操作，通过索引就可以快速地找到元素存储的地址。
	// 在数组和切片中，切片又更高效，因为它在赋值、函数传参的时候，并不会把所有的元素都复制一遍，而只是复制SliceHeader的三个字段就可以了
	// 下面可以证明数组在传参过程中被复制了
	fmt.Printf("%p\n", &a1)
	printArrayInfo(a1)
	fmt.Println()

	// 切片的高效还体现在for range循环中
	// 切片基于指针的封装是它效率高的根本原因，因为可以减少内存的占用，以及减少内存复制时的时间消耗。

	// --------------------  string和 []byte互转 --------------------
	// Go语言通过先分配一个内存再复制内容的方式，实现string和[]byte之间的强制转换
	s3 := "张三"
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)))
	b3 := []byte(s3)
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b3)))
	s4 := string(b3)
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)))
	fmt.Println()

	// StringHeader和SliceHeader一样，代表的是字符串在程序运行时的真实结构
	// []byte(s)和string(b)这种强制转换会重新拷贝一份字符串，如果字符串非常大，内存开销大

	// []byte转string，就等于通过unsafe.Pointer把*SliceHeader转为*StringHeader，也就是*[]byte转*string
	// 没有申请新内存（零拷贝）
	b5 := []byte{229, 188, 160, 228, 184, 137}
	s5 := *(*string)(unsafe.Pointer(&b5))
	fmt.Println(s5)
	fmt.Printf("%#v %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b5)), (*reflect.StringHeader)(unsafe.Pointer(&s5)))
	fmt.Println()

	// SliceHeader有Data、Len、Cap三个字段，StringHeader有Data、Len两个字段
	// 反过来却不行了，因为*StringHeader缺少*SliceHeader所需的Cap字段，需要我们自己补上一个默认值。
	s6 := "张三"
	b6h := (*reflect.SliceHeader)(unsafe.Pointer(&s6))
	b6h.Cap = b6h.Len
	b6 := *(*[]byte)(unsafe.Pointer(b6h))
	fmt.Println(b6)
	fmt.Printf("%#v %#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s6)), b6h)
	// 通过unsafe.Pointer把string转为[]byte后，不能对[]byte进行修改，比如不可以进行b1[0]=12这种操作，会报异常，导致程序崩溃。这是因为在Go语言中string内存是只读的。
	// b6[0] = 1 报错

	// 标准库里，strings.Builder这个结构体，它内部有buf字段存储内容，在通过String方法把[]byte类型的buf转为string的时候，就使用unsafe.Pointer提高了效率
}
