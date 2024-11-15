package main

import "fmt"

func main() {
	// Activity 1
	type Currency struct {
		Name string
		Rate float32
	}

	currencies := map[string]Currency{
		"USD": {"US dollar", 1.1318},
		"JPY": {"Japanese yen", 121.05},
		"GBP": {"Pound sterling", 0.90630},
		"CNY": {"Chinese yuan renminbi", 7.9944},
		"SGD": {"Singapore dollar", 1.5743},
		"MYR": {"Malaysian ringgit", 4.8390},
	}

	fmt.Println("Currencies: ")
	for key, value := range currencies {
		fmt.Printf("%s (%s): %f\n", value.Name, key, value.Rate)
	}

	// Activity 2:
	var curr1 string
	var curr2 string
	var amount float32

	fmt.Println("Enter the currency code to convert from (e.g., USD, JPY): ")
	_, err := fmt.Scanln(&curr1)
	if err != nil {
		fmt.Println("Error: Invalid input. Please input a string")
		return
	}

	fromCurrencyStruct, fromExists := currencies[curr1]
	if !fromExists {
		fmt.Println("Invalid currency code for 'from' currency. Please try again.")
		return
	}

	fmt.Println("How much do you want to convert? ")
	_, err = fmt.Scanln(&amount)
	if err != nil || amount <= 0 {
		fmt.Println("Error: Invalid number. Please enter a valid number value")
		return
	}

	fmt.Println("Enter the currency code to convert to (e.g., USD, JPY): ")
	_, err = fmt.Scanln(&curr2)
	if err != nil {
		fmt.Println("Error: Invalid input. Please input a string")
		return
	}

	toCurrencyStruct, toExists := currencies[curr2]
	if !toExists {
		fmt.Println("Invalid currency code for 'from' currency. Please try again.")
		return
	}

	// Perform the conversion using the formula
	convertedAmount := amount * (toCurrencyStruct.Rate / fromCurrencyStruct.Rate)

	// Print the result
	fmt.Printf("%.2f %s is equal to %.2f %s.\n", amount, curr1, convertedAmount, curr2)
}
