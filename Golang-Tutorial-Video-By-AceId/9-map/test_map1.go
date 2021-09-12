// map 的声明
package main

import "fmt"

func main() {
	// 声明一个 map，key 是 string 类型，value 是 int 类型
	var myMap1 map[string]int
	// 只声明了的 myMap1 == nil 返回 true
	fmt.Printf("myMap1: %v, is nil: %t\n", myMap1, myMap1 == nil)

	// 声明一个 map 并设置 size = 4
	myMap2 := make(map[string]int, 4)
	myMap2["one"] = 1
	myMap2["two"] = 2
	myMap2["three"] = 3

	// map 中元素是无序的
	fmt.Printf("myMap2: %v, len: %d\n", myMap2, len(myMap2))

	// 用 make 声明一个 map 但不设置 size
	// 注意它不是一个 nil map
	myMap3 := make(map[int]string)
	fmt.Printf("myMap3: %v, is nil: %t, len: %d\n", myMap3, myMap3 == nil, len(myMap3))
	myMap3[1] = "one"
	myMap3[2] = "two"
	myMap3[3] = "three"
	fmt.Printf("myMap3: %v, is nil: %t, len: %d\n", myMap3, myMap3 == nil, len(myMap3))

	// 直接声明 map 并带有初始值
	myMap4 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Printf("myMap4: %v\n", myMap4)
}
