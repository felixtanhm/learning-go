package main

import "fmt"

func main() {
	var start, end int

	// Ask the user for two values
	fmt.Print("Enter the starting value: ")
	fmt.Scanln(&start)
	fmt.Print("Enter the ending value: ")
	fmt.Scanln(&end)

	// Count up from start to end
	fmt.Printf("\nCounting up from %d to %d:\n", start, end)
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}

	// Count down from end to start
	fmt.Printf("\nCounting down from %d to %d:\n", end, start)
	for i := end; i >= start; i-- {
		fmt.Println(i)
	}
}
