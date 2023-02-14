package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	// ResponseWriter的Write方法返回成功写入的字节数，以及遇到的任何错误
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

// 服务器向viewHandler传递一个http.ResponseWriter，用于向浏览器响应写入数据
// 以及一个指向http.Request值的指针，该值表示浏览器的请求。
func aHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, A")
}

func bHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, B")
}

func cHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, C")
}

func main() {
	http.HandleFunc("/a", aHandler)
	http.HandleFunc("/b", bHandler)
	http.HandleFunc("/c", cHandler)
	// http.ListenAndServe，它启动Web服务器
	// 第二个参数中的nil值只表示将使用通过HandleFunc设置的函数来处理请求。
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}
