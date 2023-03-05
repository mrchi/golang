package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	Name string
	Age  int
}

func main() {
	// unsafe是不安全的，尽可能不使用它。
	// 虽然它不安全，但它也有优势，那就是可以绕过Go的内存安全机制，直接对内存进行读写。

	// 不能直接进行指针类型转换
	// ip := new(int)
	// 这句会报错 cannot convert ip (variable of type *int) to type *float64
	// var fp *float64 = (*float64)(ip)

	// -------------------- unsafe.Pointer --------------------

	// unsafe.Pointer是一种特殊意义的指针，可以表示任意类型的地址
	// 通过unsafe.Pointer做中转，可以进行指针类型转换
	i := 10
	ip := &i
	fmt.Printf("%#v, %T\n", ip, ip)

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	fmt.Printf("%#v, %T\n", fp, fp)
	*fp *= 3
	fmt.Printf("%#v, %T\n", i, i)
	fmt.Println()

	// 底层实现中，type Pointer *ArbitraryType
	// 按Go语言官方的注释，ArbitraryType可以表示任何类型
	// 则 Pointer 是一个通用型的指针，足以表示任何内存地址

	// -------------------- uintptr --------------------

	// uintptr也是一种指针类型，它足够大，可以表示任何指针
	// unsafe.Pointer不能进行运算，比如不支持+（加号）运算符操作，但是uintptr可以
	// 通过它，可以对指针偏移进行计算，这样就可以访问特定的内存，达到对特定内存读写的目的，这是真正内存级别的操作。

	p := new(person)
	fmt.Printf("%#v\n", *p)
	// 把*person类型的指针变量p通过unsafe.Pointer转换为*string类型的指针变量pName。
	// person这个结构体的第一个字段就是string类型的Name，所以pName这个指针就指向Name字段（偏移为0）
	pName := (*string)(unsafe.Pointer(p))
	*pName = "张三"
	// 先要把指针变量p通过unsafe.Pointer转换为uintptr，这样才能进行地址运算。
	// 偏移量可以通过函数unsafe.Offsetof计算出来，该函数返回的是一个uintptr类型的偏移量
	// 最后通过unsafe.Pointer转换后的*int类型的指针变量pAge。(*int)强类型转换，转换后才能对这块内存进行赋值操作。
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Printf("%#v\n", *p)
	fmt.Println()

	// 指针运算的核心在于它操作的是一个个内存地址，通过内存地址的增减，就可以指向一块块不同的内存并对其进行操作，而且不必知道这块内存被起了什么名字（变量名）

	// 指针转换规则
	// 1)任何类型的*T都可以转换为unsafe.Pointer。
	// 2)unsafe.Pointer也可以转换为任何类型的*T。
	// 3)unsafe.Pointer可以转换为uintptr。
	// 4)uintptr也可以转换为unsafe.Pointer。
	// unsafe.Pointer主要用于指针类型的转换，而且是各个指针类型转换的桥梁。
	// uintptr主要用于指针运算，尤其是通过偏移量定位不同的内存。

	// -------------------- unsafe.Sizeof --------------------
	// Sizeof函数可以返回一个类型所占用的内存大小，这个大小只与类型有关，与类型对应的变量存储的内容大小无关
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(1)), unsafe.Sizeof(int16(1)), unsafe.Sizeof(int32(1)))
	fmt.Println(unsafe.Sizeof(float32(1)), unsafe.Sizeof(float64(1)))
	fmt.Println(unsafe.Sizeof(string("张三")))
	fmt.Println(unsafe.Sizeof([]string{"张三", "李四"}))
	fmt.Println()

	// 与平台有关的int类型，要看平台是32位还是64位，会取最大的
	fmt.Println(unsafe.Sizeof(int(1)))
	fmt.Println()

	// 一个结构体的内存占用大小，等于它包含的字段类型内存占用大小之和。
	fmt.Println(unsafe.Sizeof(person{"张三", 22}))
	fmt.Println()
}
