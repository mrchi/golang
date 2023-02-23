package main

import "fmt"

type Person struct {
	name string
	age  uint
}

// 接口的实现并没有通过任何关键字（比如Java中的implements），所以Go语言的接口是隐式实现的。
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

func (p *Person) StringP() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

// ----------------------------------------

func PrintString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// ----------------------------------------

type StringPer interface {
	StringP() string
}

func PrintStringP(s StringPer) {
	fmt.Println(s.StringP())
}

// ----------------------------------------

func NewPerson(name string) *Person {
	return &Person{name: name}
}

// ----------------------------------------

func NewStringer(name string) fmt.Stringer {
	return &Person{name: name}
}

// ----------------------------------------
type Student struct {
	id string
	Person
}

func (s Student) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Age: %d", s.id, s.name, s.age)
}

func main() {
	// 采用字面量初始化结构体时，初始化值的顺序很重要，必须与字段定义的顺序一致。
	p := Person{"张三", 22}
	fmt.Printf("%#v\n", p)
	fmt.Println()

	// 通过方法表达式调用方法，第一个参数必须是接收者，然后才是方法自身的参数。
	// 注意方法表达式只能调用「接收者是值类型变量」的方法，「接收者是指针类型变量」的方法无法调用
	fmt.Println(Person.String(p))
	fmt.Println()

	// 以值类型接收者实现接口的时候，不管是类型本身，还是该类型的指针类型，都实现了该接口，例如上面的 Person.String
	PrintString(p)
	PrintString(&p)
	// 以指针类型接收者实现接口的时候，只有对应的指针类型才被认为实现了该接口。例如上面的 Person.StringP
	PrintStringP(&p)
	fmt.Println()

	// 工厂函数一般用于创建自定义的结构体，便于使用者调用
	// 通过工厂函数创建自定义结构体的方式，可以让调用者不用太关注结构体内部的字段，只需要给工厂函数传参就可以了。
	fmt.Printf("%#v\n", *NewPerson("李四"))
	fmt.Println()

	// 工厂函数也可以用来创建一个接口，它的好处就是可以隐藏内部具体类型的实现，让调用者只需关注接口的使用即可。
	i := NewStringer("王五")
	fmt.Println(i.String())
	fmt.Println()

	// 如果外部类型定义了与内部类型同样的方法，那么外部类型会覆盖内部类型，这就是方法的覆写
	s := Student{"0001", Person{"马六", 3}}
	fmt.Println(s.String())
}
