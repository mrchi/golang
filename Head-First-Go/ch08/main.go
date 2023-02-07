package main

import (
	"fmt"
	"magazine"
)

// 基于 struct 定义类型
type car struct {
	name     string
	topSpeed float64
}

func main() {
	// 定义 struct，赋值
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	myStruct.number = 3.14
	fmt.Printf("%T\n", myStruct)
	fmt.Printf("%v\n", myStruct)
	fmt.Printf("%#v\n", myStruct)
	fmt.Println()

	// 基于自定义类型创建变量
	var volvoSUV car
	volvoSUV.name = "Volvo"
	volvoSUV.topSpeed = 178.9
	fmt.Printf("%T\n", volvoSUV)
	fmt.Printf("%v\n", volvoSUV)
	fmt.Printf("%#v\n", volvoSUV)
	fmt.Println()

	// 自定义类型作为函数参数
	showInfo(volvoSUV)
	fmt.Println()

	// 自定义类型的指针作为函数参数
	// 为了节省内存空间和效率，应该尽量在函数中只传递一个struct指针的时候，保证内存中只有一个原始的struct。
	modifyTopSpeed(&volvoSUV)
	showInfo(volvoSUV)
	fmt.Println()

	// struct 字面量
	subscriber := magazine.Subscriber{Name: "Zhang3", Rate: 4.99, Active: true}
	fmt.Printf("%#v\n", subscriber)
	fmt.Println()

	// 给内部 struct 赋值
	address := magazine.Address{Street: "Xueyuan Road", City: "Beijing"}
	subscriber.HomeAddress = address
	fmt.Printf("%#v\n", subscriber)
	subscriber.HomeAddress.PostalCode = "100083"
	fmt.Printf("%#v\n", subscriber)
	fmt.Println()

	// 嵌入匿名 struct
	employee := magazine.Employee{Name: "Li4", Salary: 99.99}
	employee.PostalCode = "100083"
	employee.Address.City = "Beijing"
	fmt.Printf("%#v\n", employee)
}

func showInfo(c car) {
	fmt.Println("Name of car is:", c.name)
	fmt.Println("TopSpeed of car is:", c.topSpeed)
}

func modifyTopSpeed(c *car) {
	// 使用点运算符在 struct 指针和 struct 上都可访问字段。
	// c.topSpeed 等同于 (*c).topSpeed
	fmt.Println("Before change:", (*c).topSpeed)
	c.topSpeed = 300
}
