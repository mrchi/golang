// 变量的 pair 结构
package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct{}

func (this *Book) ReadBook() {
	fmt.Println("ReadBook()...")
}

func (this *Book) WriteBook() {
	fmt.Println("WriteBook()...")
}

func main() {
	// b 是一个 Book 实例的指针
	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	// 强制类型转换，而能够成功的前提是，都是 Book 类型，Book 类型实现了两个 interface 需要的方法
	w = r.(Writer)
	w.WriteBook()
}
