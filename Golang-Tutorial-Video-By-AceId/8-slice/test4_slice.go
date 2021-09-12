// slice 的长度和容量
package main

import "fmt"

func main() {
	// 生成长度为 3，容量为 5 的 slice
	numbers := make([]int, 3, 4)
	fmt.Printf("slice = %v, length = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	// 追加元素，长度加 1，容量未满，容量不变
	numbers = append(numbers, 1)
	fmt.Printf("slice = %v, length = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))

	// 再次追加，容量已满，容量自动扩容（2 倍）
	numbers = append(numbers, 2)
	fmt.Printf("slice = %v, length = %d, capacity = %d\n", numbers, len(numbers), cap(numbers))
}
