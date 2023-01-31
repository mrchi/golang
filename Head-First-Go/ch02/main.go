package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 调用方法
	now := time.Now()
	fmt.Println(now.Year())

	// 调用方法
	broken := "Hell#, w#rld"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(replacer.Replace(broken))
}
