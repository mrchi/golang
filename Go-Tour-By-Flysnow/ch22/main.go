package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// init，通过它可以实现包级别的一些初始化操作
// init函数没有返回值，也没有参数，它先于main函数执行
// 一个包中可以有多个init函数，但是它们的执行顺序并不确定
// init函数的作用是什么呢？其实就是在导入一个包时，可以对这个包做一些必要的初始化操作
func init() {
	fmt.Println("This is init function")
}

func main() {
	r := gin.Default()
	r.Run()
}
