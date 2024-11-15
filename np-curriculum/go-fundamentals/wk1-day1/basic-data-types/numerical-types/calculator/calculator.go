package main

import "fmt"

func main() {
	var num1 int
	var operator string
	var num2 int
	fmt.Println("Enter your first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Error: Invalid integer. Please enter a valid integer value")
		return
	}
	fmt.Println("Choose your operator (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	fmt.Println("Enter your second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Error: Invalid integer. Please enter a valid integer value")
		return
	}
	if operator == "+" {
		fmt.Println(num1 + num2)
	} else if operator == "-" {
		fmt.Println(num1 - num2)
	} else if operator == "*" {
		fmt.Println(num1 * num2)
	} else if operator == "/" {
		if num2 != 0 {
			fmt.Println(num1 / num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	} else {
		fmt.Println("Invalid operator")
	}
}
