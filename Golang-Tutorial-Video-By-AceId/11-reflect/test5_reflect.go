// 复杂类型的 reflect
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Printf("User %v is called...\n", this)
}

func main() {
	user := User{1, "xiaoming", 18}
	user.Call()

	GetFieldAndMethod(user)
}

func GetFieldAndMethod(input interface{}) {
	// 获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Printf("inputType = %v\n", inputType)

	// 获取 input 的 value
	inputValue := reflect.ValueOf(input)
	fmt.Printf("inputValue = %v\n", inputValue)

	// 通过 input 的 type 获取 Field
	for i := 0; i < inputType.NumField(); i++ {
		fmt.Println("-----")

		// 获取 Type Field，类型是 reflect.StructField
		TypeField := inputType.Field(i)
		fmt.Printf("TypeField = %v, type = %T\n", TypeField, TypeField)

		// 获取 Value Field, 类型是 reflect.Value
		ValueField := inputValue.Field(i)
		fmt.Printf("ValueField = %v, type = %T\n", ValueField, ValueField)

		// 获取 Field 的 name, type 和 value
		fmt.Printf("%v: %v = %v\n", TypeField.Name, TypeField.Type, ValueField.Interface())
	}
	// ? 只能取到 func (this User) 定义的，取不到 func (this *User) 定义的
	for i := 0; i < inputType.NumMethod(); i++ {
		fmt.Println("*****")
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
