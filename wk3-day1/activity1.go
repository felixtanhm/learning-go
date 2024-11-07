package main

import (
	"fmt"
	"sync"
)

type game struct {
	name  string
	price float64
}

func main() {
	gamesList := []game{{name: "Felix", price: 12.2}, {name: "Fiona", price: 1.5}}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		manageGameItems(&gamesList, &mu)

	}()
	go func() {
		defer wg.Done()
		manageGameItems(&gamesList, &mu)

	}()
	wg.Wait()
}

func manageGameItems(arr *[]game, mu *sync.Mutex) {
	choice := 0
	fmt.Println("What would you like to do? ")
	fmt.Println("1: Display all games available")
	fmt.Println("2: Add a new game")
	fmt.Scan(&choice)

	if choice > 2 || choice < 1 {
		fmt.Println("Please input the number corresponding to the choices provided.")
	}
	if choice == 1 {
		mu.Lock()
		displayGames(*arr)
		mu.Unlock()
	}
	if choice == 2 {
		mu.Lock()
		*arr = addGame(*arr)
		mu.Unlock()
	}
}

func addGame(arr []game) []game {
	gameName := ""
	gamePrice := 0.00
	fmt.Println("What is the name of the game?")
	fmt.Scan(&gameName)
	fmt.Println("What is the price of the game?")
	fmt.Scan(&gamePrice)
	newGame := game{
		name:  gameName,
		price: gamePrice,
	}
	arr = append(arr, newGame)
	return arr
}

func displayGames(arr []game) {
	if len(arr) < 1 {
		fmt.Println("Sorry there are no games available.")
	} else {
		fmt.Println("Here are the current list of games: ")
		for i, v := range arr {
			fmt.Printf("%d. %s ($%.2f)\n", i+1, v.name, v.price)
		}
	}
}
