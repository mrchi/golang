package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	elapsed1 := time.Since(start1)
	fmt.Println(elapsed1)

	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed2 := time.Since(start2)
	fmt.Println(elapsed2)
}

// args: a b c d e f g h i j k l m n o p q r s t u v w x y z
// elapsed1 = 47.561µs
// elapsed2 = 2.406µs
