// struct 结构体
package main

import "fmt"

// 定义类型
type myint int

// 定义结构体
type Book struct {
	title  string
	author string
	price  float64
}

func main() {
	var a myint = 10
	fmt.Printf("a = %v, type = %T\n", a, a)

	var book Book
	// 属性赋值
	book.title = "The Go Programing Language"
	book.author = "Alan"
	book.price = 79.00
	fmt.Printf("book = %v, type = %T\n", book, book)

	// 结构体作为函数参数传递时，是值传递，在函数中改变值不会影响原值
	changeBook(book)
	fmt.Printf("book = %v, type = %T\n", book, book)
}

func changeBook(book Book) {
	book.price = 999.00
}
