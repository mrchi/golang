package main

import "fmt"

func main() {
	var freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, ftoc(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, ftoc(boilingF))
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
