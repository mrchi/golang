// 变量的 pair 结构
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// tty: type = *os.File, value = /dev/tty
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
	}

	// r: type = ?, value = ?
	var r io.Reader
	// r: type = *os.File, value = /dev/tty
	r = tty

	// r: type = ?, value = ?
	var w io.Writer
	// 强制转换为 io.Writer 类型
	// w: type = *os.File, value = /dev/tty
	w, _ = r.(io.Writer)
	fmt.Printf("w = %v, type = %T\n", w, w)

	w.Write([]byte("HELLO\n"))
}
