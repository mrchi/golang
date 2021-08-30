package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "STDIN", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			data := strings.Split(line, "|")
			fmt.Printf("%d\t%s\t%s\n", n, data[0], data[1])
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[filename+"|"+input.Text()]++
	}
}
