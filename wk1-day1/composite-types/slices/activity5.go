package main

import (
	"fmt"
)

func main() {
	// Create a slice of slices to store temperature readings for each room
	temperatures := [][]float64{
		{20, 21, 23, 25, 22},     // Room 1
		{27, 23, 25, 20, 30, 24}, // Room 2
		{22, 23, 24, 22},         // Room 3
	}

	// Iterate over each room and calculate the average temperature
	for i, roomTemps := range temperatures {
		var sum float64
		for _, temp := range roomTemps {
			sum += temp
		}
		// Calculate the average temperature for the room
		average := sum / float64(len(roomTemps))
		fmt.Printf("Average temperature for Room %d: %.2f°C\n", i+1, average)
	}
}
