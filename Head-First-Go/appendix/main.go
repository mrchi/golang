package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
)

func sendLetters(channel chan string) {
	for _, RuneLetter := range "中华小当家" {
		channel <- string(RuneLetter)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// %o 八进制格式化输出
	fmt.Printf("%04o\n", 8)
	fmt.Printf("%04o\n", 64)
	fmt.Println()

	// os.Create 用于创建文件（权限 0666），或在文件存在时清空文件。O_RDWR
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// if 语句的初始化语句
	// 与for循环类似，Go允许你在if语句中的条件之前添加初始化语句
	// 初始化语句中声明的变量的作用域仅限于if语句的条件表达式及其块
	// 如果一个函数有多个返回值，而你需要其中一个在if语句中，另一个在if语句外，那么你可能无法在if初始化语句中调用它
	if _, err := file.WriteString("hello 中国"); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	// os.Open() 打开文件，O_RDONLY
	file, err = os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// os.File.Read() 读取与 b 相同字节数的文件内容到 b 中
	content := []byte("abcde")
	if _, err := file.Read(content); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	fmt.Println()

	// switch...case...default 语句
	// Go会在case代码末尾自动退出switch，不需要 break
	// 如果你希望下一个case的代码也能运行，那么可以在一个case中使用fallthrough关键字。
	rand.Seed(time.Now().Unix())
	switch rand.Intn(4) + 1 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	default:
		fmt.Println("3 or 4")
	}
	fmt.Println()

	// ----------------------rune---------------------- //

	// Go使用UTF-8，这是一种表示Unicode字符的标准，每个字符使用1到4个字节
	asciiString := "ABCDE"
	utf8String := "南京大排档"
	// 当你将字符串传递给len函数时，它将返回以字节（而不是符文）为单位的长度。
	// 如果需要字符串的字符长度，则应该使用unicode/utf8包的RuneCountInString函数。
	fmt.Println(len(asciiString), utf8.RuneCountInString(asciiString))
	fmt.Println(len(utf8String), utf8.RuneCountInString(utf8String))
	fmt.Println()

	// Go使用rune类型的值来表示Unicode值
	// 要使用部分字符串，应该将它们转换为rune值的切片，而不是byte值的切片
	asciiRune := []rune(asciiString)
	utf8Rune := []rune(utf8String)
	fmt.Println(string(asciiRune[:3]), string(utf8Rune[:3]))
	fmt.Println()

	// 对字符串使用for...range循环，它一次处理一个符文
	// 你提供的第一个变量将被分配给字符串中的当前字节索引（而不是rune索引）
	for position, currentRune := range utf8String {
		fmt.Println(position, string(currentRune))
	}
	fmt.Println()

	// ----------------------有缓冲的 channel---------------------- //

	// 有缓冲的channel可以在导致发送的goroutine阻塞之前保存一定数量的值。
	fmt.Println(time.Now().Unix())

	// 通过给make传递第二个参数来创建有缓冲的channel，该参数包含channel应该能够在其缓冲区中保存的值的数量。
	channel := make(chan string, 2)

	go sendLetters(channel)
	time.Sleep(5 * time.Second)
	// 缓冲区满后额外的发送操作才会导致goroutine阻塞
	// goroutine从channel接收一个值时，它从缓冲区中提取最早添加的值。
	fmt.Println(<-channel, time.Now().Unix()) // 拿缓冲区中的“中”
	fmt.Println(<-channel, time.Now().Unix()) // 拿缓冲区中的“华”
	fmt.Println(<-channel, time.Now().Unix()) // 拿已经发送到 channel 但造成了阻塞的“小”
	fmt.Println(<-channel, time.Now().Unix()) // channel 空了，被阻塞，等待 1s 后拿到“当”
	fmt.Println(<-channel, time.Now().Unix()) // channel 空了，被阻塞，等待 1s 后拿到“家”
}
