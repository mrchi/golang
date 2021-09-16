// interface，类似 Python 中的 duck type
package main

import "fmt"

// 本质上是一个指针
type Animal interface {
	Sleep()
	GetColor() string
}

// Dog 类定义
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping...")
}

func (this *Dog) GetColor() string {
	return this.color
}

// Cat 类定义
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping...")
}

func (this *Cat) GetColor() string {
	return this.color
}

// interface 作为参数
func ShowAnimal(animal Animal) {
	animal.Sleep()
	fmt.Printf("Color = %v\n", animal.GetColor())
}

func main() {
	dog := Dog{"yellow"}
	cat := Cat{"White"}

	// 传给 interface 对象的是个指针
	ShowAnimal(&dog)
	ShowAnimal(&cat)
}
