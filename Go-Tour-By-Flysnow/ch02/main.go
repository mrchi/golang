package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	// int -> string
	countString := strconv.Itoa(3)
	fmt.Printf("%#v\n", countString)

	// string -> int
	count, err := strconv.Atoi("4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", count)

	// float -> string
	piString := strconv.FormatFloat(3.14, byte(102), 1, 64) // byte(102) 指代字母 f
	fmt.Printf("%#v\n", piString)

	// string -> float
	pi, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", pi)

	// bool -> string
	areYouString := strconv.FormatBool(false)
	fmt.Printf("%#v\n", areYouString)

	// string -> bool
	areYou, err := strconv.ParseBool("TRUE")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", areYou)

	fmt.Println()

	// 处理字符串的 strings 包
	s1 := "Hello world"
	fmt.Println(strings.Contains(s1, "Hello"))
	fmt.Println(strings.ContainsAny(s1, "Gulf"))
	fmt.Println(strings.Count(s1, "lo"))
	fmt.Println(strings.Index(s1, "ell"))
	fmt.Println(strings.HasPrefix(s1, "He"))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToTitle(s1))
}
