// 包的导入和包的 init 函数
package main

import (
	mylib1 "github.com/mrchi/golang/Golang-Tutorial-Video-By-AceId/5-init/lib1" // 导入包并设置别名，通过别名调用包中的函数
	. "github.com/mrchi/golang/Golang-Tutorial-Video-By-AceId/5-init/lib2"      // 导入包中函数到当前命名空间，可以直接调用包中的函数
	_ "github.com/mrchi/golang/Golang-Tutorial-Video-By-AceId/5-init/lib3"      // 导入包但不使用，用于只想调用包中的 init() 函数
)

func main() {
	mylib1.Lib1Test()
	Lib2Test()
}
