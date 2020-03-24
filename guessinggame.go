package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Are you a genius...Lets play!")

	//random number generator
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	//generates numbers from 0 to 10 inclusive
	secretNumber := randomizer.Intn(10)

	var guess int
	
	//prompt user for number until game is won
	for {
		fmt.Println("What's the magic number?")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("That's bigger")
		} else if guess < secretNumber {
			fmt.Println("That's smaller")
		} else {
			fmt.Println("Awesome, you got it!")
			break
		}
	}
}
