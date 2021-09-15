// OOP 继承
package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	// 继承
	Human
	age int
}

// 重写父类方法
func (this *SuperMan) Walk() {
	fmt.Println("SuperMan.Walk()...")
}

// 定义自己的方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	h := Human{"Sniper", "male"}
	h.Eat()
	h.Walk()

	// 子类定义方式 1
	s := SuperMan{Human{"Clerk", "male"}, 320}
	s.Eat() // 继承到的父类的方法
	s.Walk()
	s.Fly()

	// 子类定义方式 2
	var s2 SuperMan
	s2.name = "Clerk"
	s2.sex = "male"
	s2.age = 320
	fmt.Println(s2)
}
