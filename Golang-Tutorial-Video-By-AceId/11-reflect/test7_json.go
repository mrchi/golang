// struct 与 JSON 的转换
package main

import (
	"encoding/json"
	"fmt"
)

type Moive struct {
	// json tag 指定 JSON 中的 key
	Name   string   `json:"title"`
	Rating float64  `json:"rating"`
	year   int      `json:"year"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Moive{"肖申克的救赎", 9.7, 1994, []string{"a", "b"}}

	// json dumps，只会输出公开（大写开头）的变量
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s, %%v = %v\n", jsonStr, jsonStr)

	// json loads
	var myMovie Moive
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	}
	fmt.Printf("myMovie = %v\n", myMovie)
}
