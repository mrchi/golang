package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ScanDirectory(path string) error {
	// print 当前目录
	fmt.Println(path)

	// ls 当前目录
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	// 打印文件路径，递归子目录
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}

	return nil
}

func ReportPanic() {
	// Go提供了一个内置的recover函数，可以阻止程序陷入panic。
	// 当出现panic时，recover返回传递给panic的任何值
	p := recover()

	// 在正常程序执行过程中调用recover时，它只返回nil
	// 因此，我们要做的第一件事是测试从recover返回的panic值是否为nil
	if p == nil {
		return
	}

	// recover的返回值的类型也是interface{}，不能直接对其调用方法。
	// 要对panic值调用方法或执行其他操作，需要使用类型断言将其转换回其底层类型
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		// 恢复 panic，只需用同样的panic值再次调用panic
		panic(p)
	}
}

func Scan() {
	// 在panic期间，任何延迟的函数调用都将完成。
	// 因此，可以在一个单独的函数中放置一个recover调用，并在引发panic的代码之前使用defer调用该函数。
	defer ReportPanic()
	ScanDirectory(".")
	// 调用recover不会导致在出现panic时恢复执行，至少不会完全恢复。
	// 产生panic的函数将立即返回，而该函数块中 panic 之后的任何代码都不会执行。
	fmt.Println("I won't run.")

}

func main() {
	Scan()
	// 在产生panic的函数返回之后，正常的执行将恢复。
	fmt.Println("Over")
}
