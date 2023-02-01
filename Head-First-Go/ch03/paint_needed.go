package main

import (
	"fmt"
	"log"
)

var metersPerLiter float64

func main() {
	metersPerLiter = 10.0

	needed, err := paintNeeded(3.123, 4.4543)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed.\n", needed)

	needed, err = paintNeeded(3.123, -4.4543)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed.\n", needed)

}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	return width * height / metersPerLiter, nil
}
