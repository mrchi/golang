package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	errorCode int
	errorMsg  string
}

func (m *MyError) Error() string {
	return fmt.Sprintf("Code=%d, Msg=%s", m.errorCode, m.errorMsg)
}

func GetOne(input int) (int, error) {
	if input == 1 {
		return 1, nil
	} else {
		return 0, &MyError{001, "It's not 1."}
	}
}

func generatePanic() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	panic("Yahaha")
}

func main() {
	// 自定义error其实就是先自定义一个新类型，比如结构体，然后让这个类型实现error接口
	_, err := GetOne(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// error 断言
	if m, ok := err.(*MyError); ok {
		fmt.Printf("%#v\n", *m)
	}
	fmt.Println()

	// error 嵌套
	original := &MyError{002, "It's not 2."}
	wrapped := fmt.Errorf("wrapped [%w]", original)
	// 不再是 MyError 类型
	if _, ok := wrapped.(*MyError); !ok {
		fmt.Printf("%#v\n", wrapped)
		fmt.Println(wrapped.Error())
	}
	fmt.Println()

	// unwrap
	unwrapped := errors.Unwrap(wrapped)
	if _, ok := unwrapped.(*MyError); ok {
		fmt.Printf("%#v\n", unwrapped)
	}
	fmt.Println()

	// Go语言为我们提供了errors.Is函数，用来判断两个error是否是同一个
	// ，两个error相等或err包含target的情况下返回true，其他情况下返回false
	fmt.Println(errors.Is(unwrapped, original))
	fmt.Println(errors.Is(wrapped, original))
	fmt.Println(errors.Is(original, wrapped))
	fmt.Println()

	// errors.As函数，用于断言
	var myError *MyError
	fmt.Println(errors.As(wrapped, &myError))
	fmt.Println()

	// 在程序因panic异常崩溃的时候，只有被defer修饰的函数才能被执行，
	// 所以recover函数要结合defer关键字使用才能生效。
	generatePanic()
}
