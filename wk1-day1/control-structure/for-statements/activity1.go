package main

import "fmt"

func main() {
	// 1. Print numbers from 0 to 1000
	fmt.Println("Numbers from 0 to 1000:")
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
	}

	// 2. Print only the odd numbers from 0 to 1000
	fmt.Println("\nOdd numbers from 0 to 1000:")
	for i := 0; i <= 1000; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 3. Print only the even numbers from 0 to 1000
	fmt.Println("\nEven numbers from 0 to 1000:")
	for i := 0; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
