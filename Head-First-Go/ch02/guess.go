// guess challenges, players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	rand.Seed(timestamp)
	target := rand.Intn(100) + 1

	success := false
	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "left.")

		// input
		fmt.Print("Input your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// to int
		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		// compare
		if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number, it was:", target)
	}
}
