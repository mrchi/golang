// slice 的截取
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s1 := s[1:]
	fmt.Printf("s = %v, s1 = %v\n", s, s1)

	// slice 的截取是引用，改变截取的 s1 会影响原来的 s
	s1[1] = 100
	fmt.Printf("s = %v, s1 = %v\n", s, s1)

	// 若想互不影响，使用 copy 函数复制
	s2 := make([]int, 3)
	copy(s2, s)
	s2[1] = 999
	fmt.Printf("s = %v, s2 = %v\n", s, s2)
}
