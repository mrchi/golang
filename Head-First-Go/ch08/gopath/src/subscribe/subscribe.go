package main

import (
	"fmt"
	"magazine"
)

func main() {
	// struct 字面量
	subscriber := magazine.Subscriber{Name: "Zhang3", Rate: 4.99, Active: true}
	fmt.Printf("%v\n", subscriber)
	fmt.Printf("%#v\n", subscriber)

}
