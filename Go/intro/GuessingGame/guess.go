package GuessingGame

import (
	"fmt"
	"math/rand"
	"strings"
)

// Random Number
var randint = rand.Intn(100)

// Update Random No.
func NewRandomInt() int {
	randint = rand.Intn(100)

	return randint
}

func Menu() {

	var guess int = 0
	var again string
	for again != "N" {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == -1 {
			fmt.Println("Bye")
			return
		}
		if guess > randint {
			fmt.Println("Smaller")
		} else if guess < randint {
			fmt.Println("Higher")
		}
		if guess == randint {
			fmt.Println("You got It")
			fmt.Println("Would you like to play again? (Y/N) ")
			fmt.Scanln(&again)
			if strings.Compare(again, "Y") == 0 {
				randint = NewRandomInt()
				guess = 0
			} else {
				fmt.Println("Ok Bye!")
			}
		}
	}
}
