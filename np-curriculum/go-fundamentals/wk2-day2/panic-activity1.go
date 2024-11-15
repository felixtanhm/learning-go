package main

import "fmt"

func main() {
	movies := []string{"Hello World", "Harry Potter", "Dragon Ball"}
	choice := 0
	defer handlePanic()
	fmt.Println("The movies are:\n")
	for i, v := range movies {
		fmt.Printf("%v: %v\n", i+1, v)
	}
	fmt.Println("Enter the number corresponding to your favorite movie: ")
	fmt.Scan(&choice)
	fmt.Println("Oh I see, your favorite movie is", movies[choice-1])
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Invalid choice. Please select only from the choices numbers.")
	}
}
