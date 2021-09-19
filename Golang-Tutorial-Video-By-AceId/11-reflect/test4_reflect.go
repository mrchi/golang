// reflect 机制
package main

import (
	"fmt"
	"reflect"
)

func reflectNum(num interface{}) {
	// 获取变量的 type
	fmt.Printf("Type = %v\n", reflect.TypeOf(num))
	// 获取变量的 value
	fmt.Printf("Value = %v\n", reflect.ValueOf(num))
}

func main() {
	i := 16.3
	reflectNum(i)
}
