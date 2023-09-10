package main

import (
	"html/template"
	"log"
	"os"
)

var x = "<html><body>Hello {{.}}</body></html>"

func main() {
	t, err := template.New("hello").Parse(x)
	if err != nil {
		log.Fatalln(err)
	}
	// 输出到 stdout，第二个参数表示传入模板的数据，会被自动编码
	t.Execute(os.Stdout, "<h1>world</h1>")
}
