// struct tag
package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	// 用 `` 定义 tag，相当于注解
	Name string `info:"the name" remark:"whole name"`
	Sex  string `info:"the sex"`
}

func findTag(s interface{}) {
	// t 是一个 reflect.rtype 指针
	t := reflect.TypeOf(s).Elem()
	fmt.Printf("t = %v, type = %T\n", t, t)

	for i := 0; i < t.NumField(); i++ {
		// 获取 tag
		tag := t.Field(i).Tag
		fmt.Printf("info = %v, remark = %v\n", tag.Get("info"), tag.Get("remark"))
	}
}

func main() {
	var r Resume

	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.TypeOf(&r))

	// 注意这里传了个指针
	findTag(&r)
}
