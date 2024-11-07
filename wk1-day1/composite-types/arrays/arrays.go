package main

import (
	"fmt"
)

func main() {
	// Declare the array with specific size and store the operating system list
	operatingSystems := [6]string{"Windows XP", "Linux 1.0", "Raspbian Teensy", "MacOS", "IOS", "Google Android"}

	// Print out the length of each element
	fmt.Println("Activity 1: ")
	for _, os := range operatingSystems {
		fmt.Printf("'%s' has a length of: %d\n", os, len(os))
	}

	// Update specific elements in the array
	operatingSystems[0] = "Windows 10"
	operatingSystems[1] = "Linux 16.04"
	operatingSystems[2] = "Raspbian Buster"

	// Print out the updated array
	fmt.Println("")
	fmt.Println("Activity 2: ")
	fmt.Println("Updated Operating Systems List:")
	for _, os := range operatingSystems {
		fmt.Println(os)
	}

	// Convert the array to a slice
	osSlice := operatingSystems[:]

	// Append new elements to the slice
	osSlice = append(osSlice, "Ubuntu", "MS-Dos", "Solaris")

	// Print out the updated slice
	fmt.Println("")
	fmt.Println("Activity 3: ")
	fmt.Println("Updated Operating Systems List:")
	for _, os := range osSlice {
		fmt.Println(os)
	}

	// a. Print the first 3 elements
	fmt.Println("First 3 elements:")
	for i := 0; i < 3; i++ {
		fmt.Println(osSlice[i])
	}

	// b. Print the next 3 elements
	fmt.Println("\nNext 3 elements:")
	for i := 3; i < 6; i++ {
		fmt.Println(osSlice[i])
	}

	// c. Print the last 3 elements
	fmt.Println("\nLast 3 elements:")
	for i := 6; i < 9; i++ {
		fmt.Println(osSlice[i])
	}
}
