package main

import (
	"fmt"
)

func main() {
	// Array to store the temperature readings for 24 hours
	temperatures := [24]float64{
		20.1, 24.0, 27.3, 30.1, 26.4, 22.2, 20.1, 24.0, 27.3, 30.1,
		26.4, 20.1, 24.0, 27.3, 30.1, 26.4, 20.1, 24.0, 27.3, 30.1,
		26.4, 20.1, 24.0, 27.3,
	}

	// Variable to store the sum of all temperatures
	var sum float64

	// Loop through the array and sum up all temperature values
	for _, temp := range temperatures {
		sum += temp
	}

	// Calculate the average temperature
	average := sum / float64(len(temperatures))

	// Print the average temperature
	fmt.Printf("The average temperature for the day is: %.2f°C\n", average)
}
