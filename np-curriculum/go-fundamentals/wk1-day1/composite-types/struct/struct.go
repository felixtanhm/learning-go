package main

import "fmt"

// Activity #1 Declare a struct for user
type Customer struct {
	FirstName        string
	LastName         string
	Age              int
	Subscriber       bool
	HomeAddress      string
	Phone            int
	CreditAvailable  float32
	CurrentCartCost  float32
	CurrentOrderCost float32
}

func main() {
	// Activity #2 Declare 2 users using struct
	user1 := Customer{
		FirstName:        "Annakin",
		LastName:         "Skywalker",
		Age:              45,
		Subscriber:       true,
		HomeAddress:      "Death Star",
		Phone:            1234567,
		CreditAvailable:  10000.00,
		CurrentCartCost:  0.00,
		CurrentOrderCost: 0.00,
	}
	user2 := Customer{
		FirstName:        "Han",
		LastName:         "Solo",
		Age:              50,
		Subscriber:       false,
		HomeAddress:      "Tatooine",
		Phone:            4321765,
		CreditAvailable:  20000.00,
		CurrentCartCost:  0.00,
		CurrentOrderCost: 0.00,
	}
	describeCustomer(user1)
	describeCustomer(user2)
}

// Function to describe a customer in a paragraph
func describeCustomer(customer Customer) {
	subscriberStatus := "is not a subscriber"
	if customer.Subscriber {
		subscriberStatus = "is a subscriber"
	}

	fmt.Printf("%s %s, aged %d, %s who lives at %s. They have a phone number %d. %s has $%.2f in available credit, with a current cart cost of $%.2f and a current order cost of $%.2f.\n\n",
		customer.FirstName,
		customer.LastName,
		customer.Age,
		subscriberStatus,
		customer.HomeAddress,
		customer.Phone,
		customer.FirstName,
		customer.CreditAvailable,
		customer.CurrentCartCost,
		customer.CurrentOrderCost)
}
