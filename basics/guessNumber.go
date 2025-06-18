package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource((time.Now().UnixNano()))
	random := rand.New(source)
	target := random.Intn(99) + 1 //Intm(n) -> (0,n)
	fmt.Println("I've chosen number from 1 to 100. Can you guess what it is?")
	var guess int
	for {
		fmt.Print(">>> Enter your guess: ")
		fmt.Scanln(&guess) // by reference!
		if guess == target {
			fmt.Println("Yes! You win!")
			break
		}

		if guess > target {
			fmt.Println("No, your number is bigger than mine.")
		} else {
			fmt.Println("No, your number is less than mine.")
		}
	}
}
