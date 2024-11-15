package main

import "fmt"

func main() {
	fmt.Println("Activity #1: If statements")
	fmt.Println("--------------")

	// Define the numbers
	num1 := 17
	num2 := 24

	// Check which number is bigger
	if num1 > num2 {
		fmt.Printf("%d is bigger than %d\n", num1, num2)
		fmt.Printf("It is bigger by %d\n", num1-num2)
	} else if num2 > num1 {
		fmt.Printf("%d is bigger than %d\n", num2, num1)
		fmt.Printf("It is bigger by %d\n", num2-num1)
	} else {
		fmt.Println("Both numbers are equal.")
	}
	fmt.Println("")
	fmt.Println("Activity #2: If ... else if statements")
	fmt.Println("--------------")

	// Check if the number is odd or even
	if num1%2 == 0 {
		fmt.Printf("%d is an even number.\n", num1)
	} else {
		fmt.Printf("%d is an odd number.\n", num1)
	}

	// Check if the number has more than one digit
	if num1 >= 10 || num1 <= -10 {
		fmt.Printf("%d has more than one digit.\n", num1)
	} else {
		fmt.Printf("%d has only one digit.\n", num1)
	}

	fmt.Println("")
	fmt.Println("Activity #3: If ... else statements")
	fmt.Println("--------------")

	fmt.Println("Enter an integer: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Error: Invalid integer. Please enter a valid integer value")
		return
	}

	if num1%5 == 0 && num1%6 == 0 {
		fmt.Printf("%d is divisible by both 5 and 6", num1)
	} else {
		fmt.Printf("%d is not divisible by both 5 and 6", num1)
	}

	fmt.Println("")
	fmt.Println("Activity #4: If ... else if... else statements")
	fmt.Println("--------------")

	// List of authenticated usernames
	authenticatedUsers := []string{"Admin", "Robin", "John"}

	// Request for the user to input a name
	var username string
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)

	// Check the entered username using if...else if...else statements
	if username == authenticatedUsers[0] {
		fmt.Println("Welcome, Admin!")
	} else if username == authenticatedUsers[1] || username == authenticatedUsers[2] {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("Merry Men")
	}

	fmt.Println("")
	fmt.Println("Activity #5: Compound If...else statements")
	fmt.Println("--------------")
	var year int

	// Request the user to enter a year
	fmt.Print("Enter a year: ")
	_, err = fmt.Scanln(&year)

	// Check if the input was valid
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid year.")
		return
	}

	// Check if the year is a leap year
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
