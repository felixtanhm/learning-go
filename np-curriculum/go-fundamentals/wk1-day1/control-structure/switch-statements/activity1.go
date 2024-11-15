package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current day of the week
	day := time.Now().Weekday()

	// Use a switch statement to determine the day and whether it's odd or even
	switch day {
	case time.Monday:
		fmt.Println("Today is Monday.")
		fmt.Println("Monday is an odd day.")
	case time.Tuesday:
		fmt.Println("Today is Tuesday.")
		fmt.Println("Tuesday is an even day.")
	case time.Wednesday:
		fmt.Println("Today is Wednesday.")
		fmt.Println("Wednesday is an odd day.")
	case time.Thursday:
		fmt.Println("Today is Thursday.")
		fmt.Println("Thursday is an even day.")
	case time.Friday:
		fmt.Println("Today is Friday.")
		fmt.Println("Friday is an odd day.")
	case time.Saturday:
		fmt.Println("Today is Saturday.")
		fmt.Println("Saturday is an even day.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
		fmt.Println("Sunday is an odd day.")
	default:
		fmt.Println("Unknown day.")
	}
}
