// map 的使用
package main

import "fmt"

func main() {
	// 定义
	cityMap := make(map[string]string)

	// 添加元素
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["Britain"] = "London"
	fmt.Println(cityMap)

	// 删除
	delete(cityMap, "Britain")
	fmt.Println(cityMap)

	// 修改
	cityMap["Japan"] = "Beijing"
	fmt.Println(cityMap)

	// 作为参数传递时是引用传递
	printMap(cityMap)
	fmt.Println(cityMap)
}

func printMap(cityMap map[string]string) {
	// 遍历
	for key, value := range cityMap {
		fmt.Printf("key = %v, value = %v\n", key, value)
	}
	cityMap["Netherland"] = "Amsterdam"
}
