package main

import "fmt"

func main() {
	for {
		// Ask the user for a number
		var number int
		fmt.Print("Enter a number (or type 'Ctrl+C' to exit): ")
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		// Check if the number is odd or even
		if number%2 == 0 {
			fmt.Printf("%d is even.\n", number)
		} else {
			fmt.Printf("%d is odd.\n", number)
		}
	}
}
