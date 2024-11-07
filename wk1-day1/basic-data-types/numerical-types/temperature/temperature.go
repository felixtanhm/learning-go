package main

import "fmt"

func main() {
	var format int
	var temperature float64
	fmt.Println("What format is the temperature in (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit):")
	_, err := fmt.Scanln(&format)
	if err != nil || (format < 1 || format > 3) {
		fmt.Println("Error: Invalid input. Please enter 1, 2, or 3.")
		return
	}
	fmt.Print("Enter the temperature value: ")
	_, err = fmt.Scanln(&temperature)
	if err != nil {
		fmt.Println("Error: Invalid temperature. Please enter a valid number.")
		return
	}
	switch format {
	case 1:
		celsius := temperature - 273.15
		fahrenheit := (temperature-273.15)*9/5 + 32
		fmt.Printf("Celsius: %.2f°C\n", celsius)
		fmt.Printf("Fahrenheit: %.2f°F\n", fahrenheit)
	case 2:
		kelvin := temperature + 273.15
		fahrenheit := (temperature * 9 / 5) + 32
		fmt.Printf("Kelvin: %.2fK\n", kelvin)
		fmt.Printf("Fahrenheit: %.2f°F\n", fahrenheit)
	case 3:
		celsius := (temperature - 32) * 5 / 9
		kelvin := (temperature-32)*5/9 + 273.15
		fmt.Printf("Celsius: %.2f°C\n", celsius)
		fmt.Printf("Kelvin: %.2fK\n", kelvin)
	default:
		fmt.Println("Error: Invalid format selected.")
	}
}
