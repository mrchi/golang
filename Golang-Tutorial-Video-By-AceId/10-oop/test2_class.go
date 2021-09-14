// OOP 的表示与封装
package main

import "fmt"

// 定义“类”——结构体
// 结构体名大写，外部能够访问
type Hero struct {
	// 属性名大写，外部能够访问
	Name  string
	Level int
}

// 方法定义 this 要使用指针，否则为值拷贝，做修改操作时对原对象无效
// 方法名大写，外部能够访问
func (this *Hero) Show() {
	fmt.Printf("Name = %v\n", this.Name)
	fmt.Printf("Level = %v\n", this.Level)
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "IO", Level: 1}
	hero.Show()

	hero.SetName("Jugg")
	hero.Show()
}
