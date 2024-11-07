package main

import "fmt"

func main() {
	text := "The following is the account information."
	firstName := "Luke"
	lastName := "Skywalker"
	age := 20
	weight := 73.0
	height := 1.72
	remainingCredits := 123.55
	accountName := "admin"
	accountPassword := "password"
	subscribed := true

	fmt.Printf("text (type: %T) = %v\n", text, text)
	fmt.Printf("firstName (type: %T) = %v\n", firstName, firstName)
	fmt.Printf("lastName (type: %T) = %v\n", lastName, lastName)
	fmt.Printf("age (type: %T) = %v yrs old\n", age, age)
	fmt.Printf("weight (type: %T) = %v kg\n", weight, weight)
	fmt.Printf("height (type: %T) = %v m\n", height, height)
	fmt.Printf("remainingCredits (type: %T) = $%v\n", remainingCredits, remainingCredits)
	fmt.Printf("accountName (type: %T) = %v\n", accountName, accountName)
	fmt.Printf("accountPassword (type: %T) = %v\n", accountPassword, accountPassword)
	fmt.Printf("subscribed (type: %T) = %v\n", subscribed, subscribed)
}
