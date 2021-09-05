package main

import (
	"fmt"
	"os"
	"strconv"

	"./lengthconv"
	"./tempconv"
	"./weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "main: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(num)
		c := tempconv.Celsius(num)
		m := lengthconv.Metre(num)
		i := lengthconv.Inch(num)
		p := weightconv.Pound(num)
		k := weightconv.Kilogram(num)
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
		fmt.Printf("%s = %s\n", m, lengthconv.MToI(m))
		fmt.Printf("%s = %s\n", i, lengthconv.IToM(i))
		fmt.Printf("%s = %s\n", p, weightconv.PToK(p))
		fmt.Printf("%s = %s\n", k, weightconv.KToP(k))
		fmt.Println("----------")
	}
}
