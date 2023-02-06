package main

import (
	"fmt"
	"sort"
)

func main() {
	// 只声明时是 nil 类型
	var myMap1 map[string]int
	fmt.Printf("%v, %#v\n", myMap1, myMap1)

	// 使用 make 函数创建
	myMap1 = make(map[string]int)
	fmt.Printf("%v, %#v\n", myMap1, myMap1)

	// 短变量声明和赋值
	myMap2 := make(map[string]int)
	myMap2["zhang3"] = 3
	myMap2["li4"] = 4
	fmt.Printf("%v, %#v\n", myMap2, myMap2)

	// 字面量
	myMap3 := map[string]int{"zhang3": 3, "li4": 4}
	fmt.Printf("%v, %#v\n", myMap3, myMap3)

	// 读取不存在的键返回零值
	wang5 := myMap3["wang5"]
	fmt.Printf("%v\n", wang5)
	// 判断键是否存在
	wang5, ok := myMap3["wang5"]
	fmt.Printf("Value: %v, OK: %v\n", wang5, ok)
	zhang3, ok := myMap3["zhang3"]
	fmt.Printf("Value: %v, OK: %v\n", zhang3, ok)

	// 删除值
	delete(myMap3, "zhang3")
	fmt.Printf("%v\n", myMap3)
	// 删除不存在的值，不会报错
	delete(myMap3, "wang5")

	// for...range 循环，是随机顺序读取
	myMap4 := map[string]int{"zhang3": 3, "li4": 4, "wang5": 5}
	// 迭代 key 和 value
	for key, value := range myMap4 {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
	// 只迭代 key
	for key := range myMap4 {
		fmt.Printf("Key: %v\n", key)
	}
	// 只迭代 value
	for _, value := range myMap4 {
		fmt.Printf("Value: %v\n", value)
	}

	// 按顺序迭代的一个例子
	var keys []string
	for key := range myMap4 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("Key: %v, Value: %v\n", key, myMap4[key])
	}
}
