package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// for...range 循环
	var sum float64 = 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Printf("Average is %.2f\n", sum/float64(len(numbers)))
}
