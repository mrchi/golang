// 指针
package main

import "fmt"

func main() {
	var a, b int = 10, 20
	swap(&a, &b) // &a 返回 a 变量的内存地址
	fmt.Printf("a = %d, b = %d\n", a, b)

	var p *int = &a
	fmt.Println(p)

	var pp **int = &p // 二级指针，指向 a 的指针
	fmt.Println(p, *pp, &a)
}

func swap(pa *int, pb *int) {
	temp := *pa // *pa 获取 pa 指针指向的变量的值
	*pa = *pb
	*pb = temp
}
