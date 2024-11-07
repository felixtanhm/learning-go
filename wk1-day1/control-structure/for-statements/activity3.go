package main

import "fmt"

func main() {
	var start, end int

	// Ask the user for two values
	fmt.Print("Enter the starting value: ")
	fmt.Scanln(&start)
	fmt.Print("Enter the ending value: ")
	fmt.Scanln(&end)

	// Print odd numbers in the range
	fmt.Printf("\nOdd numbers between %d and %d:\n", start, end)
	for i := start; i <= end; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// Print even numbers in the range
	fmt.Printf("\nEven numbers between %d and %d:\n", start, end)
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
