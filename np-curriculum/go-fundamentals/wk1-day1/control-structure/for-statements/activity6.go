package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1

	// Variables for user input and attempts
	var guess int
	attempts := 0
	maxAttempts := 5

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("You have 5 tries to guess the correct number.")
	fmt.Println("Enter 101 at any time to end the game.")

	for attempts < maxAttempts {
		fmt.Printf("Attempt %d: Enter your guess: ", attempts+1)
		_, err := fmt.Scanln(&guess)

		// Check if the input is valid
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		// Check if the user wants to end the game
		if guess == 101 {
			fmt.Println("You have chosen to exit the game. Goodbye!")
			break
		}

		// Check if the guess is correct
		if guess == target {
			fmt.Printf("Congratulations! You guessed the correct number %d!\n", target)
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}

		// Increment the attempt counter
		attempts++

		// Check if the user has reached the maximum attempts
		if attempts == maxAttempts {
			fmt.Printf("Sorry, you've used all your attempts. The correct number was %d.\n", target)
		}
	}
}
