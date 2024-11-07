package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	defer handlePanic()
	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)
	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)
	if firstName == "" {
		panic("First name cannot be empty.")
	}
	if lastName == "" {
		panic("Last name cannot be empty.")
	}
	fmt.Printf("Hello %v %v!!! Nice to meet you!\n", firstName, lastName)
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Please make sure to enter both your first and last names.")
	}
}
