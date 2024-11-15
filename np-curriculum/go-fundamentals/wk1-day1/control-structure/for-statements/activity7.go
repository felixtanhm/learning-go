package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Checksum mapping for remainders
	checksumLetters := []rune{'A', 'Z', 'Y', 'X', 'U', 'T', 'S', 'R', 'P', 'M', 'L', 'K', 'J', 'H', 'G', 'E', 'D', 'C', 'B'}
	weights := []int{9, 4, 5, 4, 3, 2}

	// Request the user to enter a vehicle plate number
	var plate string
	fmt.Print("Enter a vehicle plate number: ")
	fmt.Scanln(&plate)
	plate = strings.ToUpper(plate)

	// Validate the plate length
	if len(plate) < 2 {
		fmt.Println("Invalid plate number. Please enter a valid vehicle plate number.")
		return
	}

	// Split the prefix (letters) and the suffix (numbers)
	prefix := plate[:len(plate)-4]
	suffix := plate[len(plate)-4 : len(plate)-1]
	checkLetter := rune(plate[len(plate)-1])

	// Process the prefix based on its length
	var prefixNums []int
	switch len(prefix) {
	case 3:
		// Use the last 2 letters
		prefixNums = append(prefixNums, letterToNumber(prefix[1]), letterToNumber(prefix[2]))
	case 2:
		// Use both letters
		prefixNums = append(prefixNums, letterToNumber(prefix[0]), letterToNumber(prefix[1]))
	case 1:
		// Use the letter for the second position, default first position to 0
		prefixNums = append(prefixNums, 0, letterToNumber(prefix[0]))
	default:
		fmt.Println("Invalid prefix length.")
		return
	}

	// Pad the suffix to 4 digits
	suffixNums := padSuffix(suffix)

	// Combine prefix and suffix numbers
	numbers := append(prefixNums, suffixNums...)

	// Multiply each number with the corresponding weight
	sum := 0
	for i, num := range numbers {
		sum += num * weights[i]
	}

	// Calculate the remainder when divided by 19
	remainder := sum % 19

	// Find the corresponding checksum letter
	expectedCheckLetter := checksumLetters[remainder]

	// Check if the provided checksum matches the calculated one
	if expectedCheckLetter == checkLetter {
		fmt.Println("The vehicle plate given is correct!")
	} else {
		fmt.Println("The vehicle plate given is not correct!")
	}
}

// Convert a letter to its corresponding number (A = 1, Z = 26)
func letterToNumber(letter byte) int {
	return int(letter-'A') + 1
}

// Pad the suffix with leading zeros to ensure it has 4 digits
func padSuffix(suffix string) []int {
	// Ensure the suffix has 4 digits
	for len(suffix) < 4 {
		suffix = "0" + suffix
	}

	// Convert each character to an integer
	var result []int
	for _, ch := range suffix {
		if unicode.IsDigit(ch) {
			result = append(result, int(ch-'0'))
		}
	}
	return result
}
