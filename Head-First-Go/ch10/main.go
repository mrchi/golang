package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	fmt.Printf("%#v, %T\n", date, date)
	fmt.Println()

	err := date.SetYear(2023)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v, %#v\n", date, date.Year())
	fmt.Println()

	err = date.SetMonth(33)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
	fmt.Printf("%#v, %#v\n", date, date.Year())
	fmt.Println()

	// 外部类型定义的方法和内部嵌入类型的方法的生存时间是一样的。
	event := calendar.Event{}
	event.SetTitle("Refresh")
	fmt.Printf("%#v, %#v\n", event, event.Title())
	fmt.Println()
}
