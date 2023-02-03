package main

import (
	"fmt"
	"greeting/chn"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Println("Language is", chn.Language)
	chn.HelloWorld()

	fmt.Print("Input a number: ")
	number, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number is %.2f\n", number)
}
