package main

import (
	"fmt"
	"math"
)

func main() {
	var sideA, sideB, angle float64

	// Prompt user for input
	fmt.Println("Enter the length of the first side (A):")
	fmt.Scanln(&sideA)

	fmt.Println("Enter the length of the second side (B):")
	fmt.Scanln(&sideB)

	fmt.Println("Enter the angle between them in degrees:")
	fmt.Scanln(&angle)

	// Convert angle from degrees to radians
	angleInRadians := angle * (math.Pi / 180)

	// Apply the cosine rule to calculate the length of the third side (C)
	sideC := math.Sqrt(math.Pow(sideA, 2) + math.Pow(sideB, 2) - 2*sideA*sideB*math.Cos(angleInRadians))

	// Print the resultant length
	fmt.Printf("The length of the third side is: %.2f\n", sideC)
}
