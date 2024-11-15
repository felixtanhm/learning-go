package main

import "fmt"

func main() {
	// Variables to store the number of each type of coin
	var oneDollar, fiftyCents, twentyCents, tenCents, fiveCents int

	// Prompt user for input
	fmt.Println("Enter the number of 1-dollar coins:")
	fmt.Scanln(&oneDollar)

	fmt.Println("Enter the number of 50-cent coins:")
	fmt.Scanln(&fiftyCents)

	fmt.Println("Enter the number of 20-cent coins:")
	fmt.Scanln(&twentyCents)

	fmt.Println("Enter the number of 10-cent coins:")
	fmt.Scanln(&tenCents)

	fmt.Println("Enter the number of 5-cent coins:")
	fmt.Scanln(&fiveCents)

	// Calculate the total amount in cents
	totalCents := (oneDollar * 100) + (fiftyCents * 50) + (twentyCents * 20) + (tenCents * 10) + (fiveCents * 5)

	// Convert total cents to dollars
	totalDollars := float64(totalCents) / 100.0

	// Calculate the number of 2-dollar notes that can be given
	twoDollarNotes := totalCents / 200
	remainingCents := totalCents % 200
	remainingDollars := float64(remainingCents) / 100.0

	// Output the results
	fmt.Printf("Total amount: $%.2f\n", totalDollars)
	fmt.Printf("2-dollar notes: %d\n", twoDollarNotes)
	fmt.Printf("Remaining change: $%.2f\n", remainingDollars)
}
