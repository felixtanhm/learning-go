package main

import (
	"fmt"
)

func main() {
	var guess int
	answer := 88
	fmt.Println("Enter an integer value: ")
	_, err := fmt.Scanln(&guess)
	if err != nil {
		fmt.Println("Error: Invalid integer. Please enter a valid integer value")
		return
	}

	if guess > answer {
		fmt.Println("Too high, try again next time")
	} else if guess < answer {
		fmt.Println("Too low, try again next time")
	} else {
		fmt.Println("Well Done! Your guess is correct")
	}
}
