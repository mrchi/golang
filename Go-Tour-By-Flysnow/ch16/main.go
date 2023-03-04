package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p person) Print(prefix string) {
	fmt.Printf("%s: Name is %s, Age is %d\n", prefix, p.Name, p.Age)
}

// struct tag，在字段后面通过反引号把一个键值对包住
// 结构体的字段可以有多个tag，用于不同的场景。如果有多个tag，要使用空格分隔。
// 完全可以把struct tag当成结构体中字段的元数据配置，使用它来做你想做的任何事情，比如orm映射、xml转换、生成swagger文档等。
type personWithTag struct {
	Name string `json:"name" bson:"person_name"`
	Age  int    `json:"age"` // json作为Key，是Go语言自带的json包解析JSON的一种约定，它会通过json这个Key找到对应的值。
}

func main() {
	// 反射，一种可以在运行时操作任意类型对象的能力
	// 在Go语言的反射定义中，任何接口都由两部分组成：接口的具体类型，以及具体类型对应的值。

	// 标准库为我们提供了两种类型reflect.Value和reflect.Type来分别表示变量的值和类型
	// 提供了两个函数reflect.ValueOf和reflect.TypeOf分别获取任意对象的reflect.Value和reflect.Type。
	i := 3
	iv := reflect.ValueOf(i)
	fmt.Printf("%T, %v, %#v\n", iv, iv, iv)
	it := reflect.TypeOf(i)
	fmt.Printf("%T, %v, %#v\n", it, it, it)
	fmt.Println()

	// ------------------- reflect.Value -------------------
	// reflect.Value 是一个 struct
	// 逆向转回原来的类型，reflect.Value为我们提供了Interface方法
	i1 := iv.Interface().(int)
	fmt.Printf("%T, %v, %#v\n", i1, i1, i1)
	fmt.Println()

	// 修改对应的值。总结一下通过反射修改一个值的规则。
	// 1)可被寻址，通俗地讲就是要向reflect.ValueOf函数传递一个指针作为参数。
	// 2)如果要修改结构体字段值的话，该字段需要是可导出的，而不是私有的，也就是该字段的首字母为大写。
	// 3)记得使用Elem方法获得指针指向的值，这样才能调用Set系列方法进行修改。
	p := person{"张三", 23}
	ppv := reflect.ValueOf(&p)
	fmt.Printf("%T, %v, %#v\n", ppv, ppv, ppv)
	fmt.Printf("%T, %v, %#v\n", ppv.Elem(), ppv.Elem(), ppv.Elem())
	fmt.Printf("%T, %v, %#v\n", ppv.Elem().FieldByName("Age"), ppv.Elem().FieldByName("Age"), ppv.Elem().FieldByName("Age"))

	ppv.Elem().FieldByName("Age").SetInt(21)
	fmt.Printf("%#v\n", p)
	fmt.Println()

	// 获取对应的底层类型, 通过Kind方法返回一个Kind类型的值，它是一个常量
	pv := reflect.ValueOf(p)
	fmt.Printf("%T, %v, %#v\n", pv.Kind(), pv.Kind(), pv.Kind())
	ppv = reflect.ValueOf(&p)
	fmt.Printf("%T, %v, %#v\n", ppv.Kind(), ppv.Kind(), ppv.Kind())
	fmt.Println()

	// ------------------- reflect.Type -------------------
	// reflect.Type是一个接口
	// 1)Implements方法用于判断是否实现了接口u。
	// 2)AssignableTo方法用于判断是否可以赋值给类型u，其实就是是否可以使用“=”，即赋值运算符
	// 3)ConvertibleTo方法用于判断是否可以转换成类型u，其实就是是否可以进行类型转换。
	// 4)Comparable方法用于判断该类型是否是可比较的，其实就是是否可以使用关系运算符进行比较

	// 遍历结构体的字段和方法
	// NumField方法获取结构体字段的数量，通过Field方法就可以遍历结构体的字段
	// 遍历结构体的方法也是同样的思路
	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("Field", pt.Field(i).Name)
	}
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("Method", pt.Method(i).Name)
	}
	fmt.Println()

	// 通过FieldByName方法获取指定的字段，通过MethodByName方法获取指定的方法
	ageField, ok := pt.FieldByName("Age")
	fmt.Printf("%T, %v, %#v\n", ageField, ageField, ageField)
	fmt.Println(ok)
	fmt.Println()

	stringMethod, ok := pt.MethodByName("String")
	fmt.Printf("%T, %v, %#v\n", stringMethod, stringMethod, stringMethod)
	fmt.Println(ok)
	fmt.Println()

	notExistField, ok := pt.FieldByName("Address")
	fmt.Printf("%T, %v, %#v\n", notExistField, notExistField, notExistField)
	fmt.Println(ok)
	fmt.Println()

	// 通过Implements方法来判断是否实现某接口
	// 尽可能通过类型断言的方式判断是否实现了某接口，而不是通过反射。
	fmt.Printf("%T, %v, %#v\n", (*fmt.Stringer)(nil), (*fmt.Stringer)(nil), (*fmt.Stringer)(nil))
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println(pt.Implements(stringerType), pt.Implements(writerType))
	fmt.Println()

	// ---------- 字符串和结构体的互转 ----------
	// json.Marshal函数，你可以把一个struct转为JSON字符串
	respJSON, err := json.Marshal(p) // 返回的是 []byte
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", string(respJSON))
	// 通过json.Unmarshal函数，你可以把一个JSON字符串转为struct。
	var p2 person
	json.Unmarshal(respJSON, &p2)
	fmt.Printf("%#v\n", p2)
	fmt.Println()

	// struct tag 是一个添加在struct字段上的标记，使用它进行辅助，可以完成一些额外的操作
	tagRespJSON, _ := json.Marshal(personWithTag{"李四", 22})
	fmt.Printf("%#v\n", string(tagRespJSON))
	var p3 personWithTag
	json.Unmarshal(tagRespJSON, &p3)
	fmt.Printf("%#v\n", p3)
	fmt.Println()

	// 如果 JSON field 和 struct 的 field 只是大小写不同，也可以不使用 tag 直接完成 JSON -> struct 的转换
	json.Unmarshal(tagRespJSON, &p2)
	fmt.Printf("%#v\n", p2)
	fmt.Println()

	// json 包是通过反射获得的 tag，要获得Key为json的tag，只需要调用sf.Tag.Get("json")即可
	p3t := reflect.TypeOf(p3)
	for i := 0; i < p3t.NumField(); i++ {
		field := p3t.Field(i)
		fmt.Println(field.Name, field.Tag, field.Tag.Get("json"))
	}
	fmt.Println()

	// 反射灵活、强大，但也存在不安全因素。它可以绕过编译器的很多静态检查，如果过多使用便会造成混乱。
	// 1)任何接口值interface{}都可以反射出反射对象，也就是ref lect.Value和ref lect.Type通过函数ref lect.ValueOf和ref lect.TypeOf获得。
	// 2)反射对象也可以还原为interface{}变量，也就是第1条定律的可逆性，通过ref lect.Value结构体的Interface方法获得。
	// 3)要修改反射的对象，该值必须可设置，也就是可寻址，可以参考上一章中修改变量的值那一节的内容来理解。

	// 用反射的方式调用方法。
	// 1) 首先要通过MethodByName方法找到相应的方法
	// 2) 然后声明参数，它的类型是[]reflect.Value
	// 3) 最后就可以通过Call方法反射调用Print方法了
	mPrint := pv.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("你好")}
	mPrint.Call(args)
}
