package main

import (
	"fmt"
)

func main() {
	// Create a slice for this month’s spendings with a capacity of 5
	spendings := make([]float64, 4, 5)
	spendings[0] = 9.50
	spendings[1] = 8.00
	spendings[2] = 10.20
	spendings[3] = 7.40

	// Print out the length and capacity of the slice
	fmt.Println("Activity 1: ")
	fmt.Printf("Length of the slice: %d\n", len(spendings))
	fmt.Printf("Capacity of the slice: %d\n", cap(spendings))

	fmt.Println("")
	fmt.Println("Activity 2: ")
	fmt.Printf("Spending for week 3: %f\n", spendings[2])
	spendings[2] = 9.80
	for i := range spendings {
		fmt.Println(spendings[i])
	}

	fmt.Println("")
	fmt.Println("Activity 3: ")
	spendings = append(spendings, 8.40, 9.40, 7.20)
	fmt.Printf("Length of the slice: %d\n", len(spendings))
	fmt.Printf("Capacity of the slice: %d\n", cap(spendings))

	fmt.Println("")
	fmt.Println("Activity 4: ")
	spendings = spendings[3:]
	fmt.Printf("Length of the slice: %d\n", len(spendings))
	fmt.Printf("Capacity of the slice: %d\n", cap(spendings))
}
