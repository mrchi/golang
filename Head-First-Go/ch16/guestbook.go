package main

import (
	"bufio"
	"fmt"

	// html/template包自动“转义”了HTML，用代码替换了导致将其视为HTML的字符
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	Signatures     []string
	SignatureCount int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)

	// 把从os.Open获得的任何error值传递给os.IsNotExist函数
	// 错误表示文件不存在，则返回true。
	if os.IsNotExist(err) {
		return nil
	}

	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	guestbook := Guestbook{
		Signatures:     signatures,
		SignatureCount: len(signatures),
	}

	// template.ParseFiles 函数加载模版文件
	// ParseFiles将返回一个指向此Template的指针，也可能返回一个error值
	//
	html, err := template.ParseFiles("view.html")
	check(err)

	// 要从Template值中获得输出，我们使用两个参数来调用其Execute方法
	// 第一个参数要求是 io.Writer 接口类型（必须具有 Write() 方法）
	err = html.Execute(writer, guestbook)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	// 对http.Request调用FormValue方法，返回指定表单字段的字符串
	signature := request.FormValue("signature")

	// flag 参数指定模式为：只写、追加、如果不存在则新建
	// 创建文件时使用 perm 参数指定的权限
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(600))
	check(err)

	// 写入文件
	_, err = fmt.Fprintln(file, signature)
	check(err)

	// 关闭文件
	err = file.Close()
	check(err)

	http.Redirect(writer, request, "/", 302)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	// ListenAndServe总是返回一个错误。（如果没有错误，ListenAndServe将永远不会返回。）
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}
