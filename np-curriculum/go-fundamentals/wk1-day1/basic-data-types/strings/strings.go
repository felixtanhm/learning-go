package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Define the lists
	list1 := [4]string{"ans", "wer", "is", "of"}
	list2 := [4]string{"-", "+", "*", "/"}
	list3 := [4]string{"5", "10", "4", "0"}

	// Extract the necessary elements from the lists
	part1 := list1[0] + list1[1] // "ans" + "wer" = "answer"
	part2 := list1[3]            // "of"
	number1 := list3[0]          // "5"
	operator := list2[1]         // "+"
	number2 := list3[2]          // "4"
	part3 := list1[2]            // "is"

	// Convert numbers to integers for calculation
	num1, _ := strconv.Atoi(number1)
	num2, _ := strconv.Atoi(number2)

	// Perform the calculation (5 + 4)
	result := num1 + num2

	// Construct the final statement
	finalStatement := fmt.Sprintf("%s %s %s %s %s %s %d",
		strings.Title(part1), part2, number1, operator, number2, part3, result)

	// Display the final statement
	fmt.Println(finalStatement)
}
