package main

import (
	"fmt"
)

func main() {
	var weight, height float64

	// Request the user to input their weight and height
	fmt.Print("Enter your weight in kilograms: ")
	_, errWeight := fmt.Scanln(&weight)
	if errWeight != nil || weight <= 0 {
		fmt.Println("Invalid weight input. Please enter a valid number.")
		return
	}

	fmt.Print("Enter your height in meters: ")
	_, errHeight := fmt.Scanln(&height)
	if errHeight != nil || height <= 0 {
		fmt.Println("Invalid height input. Please enter a valid number.")
		return
	}

	// Calculate the BMI
	bmi := weight / (height * height)
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	// Determine BMI category using a switch statement
	switch {
	case bmi < 18.5:
		fmt.Println("You are underweight.")
	case bmi >= 18.5 && bmi <= 24.9:
		fmt.Println("You have a healthy weight.")
	case bmi >= 25 && bmi <= 29.9:
		fmt.Println("You are overweight.")
	case bmi >= 30 && bmi <= 34.9:
		fmt.Println("You are obese.")
	case bmi >= 35 && bmi <= 39.9:
		fmt.Println("You are severely obese.")
	case bmi >= 40:
		fmt.Println("You are morbidly obese.")
	default:
		fmt.Println("Invalid BMI value.")
	}
}
