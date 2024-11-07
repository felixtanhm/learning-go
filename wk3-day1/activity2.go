package main

import (
	"fmt"
	"sync"
)

type bankBalance struct {
	amount   float64
	currency string
}

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	myBB := bankBalance{
		amount:   100.23,
		currency: "SGD",
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		processBankBalance(&myBB, &mu)
	}()
	go func() {
		defer wg.Done()
		processBankBalance(&myBB, &mu)
	}()
	wg.Wait()

}

func processBankBalance(bb *bankBalance, mu *sync.Mutex) {
	choice := 0
	fmt.Println("What would you like to do? ")
	fmt.Println("1: Deposit into my bank balance")
	fmt.Println("2: Withdraw from my bank balance")
	fmt.Println("3: Display my bank balance")

	fmt.Scan(&choice)

	if choice > 3 || choice < 1 {
		fmt.Println("Please input the number corresponding to the choices provided.")
	}
	if choice == 1 {
		amount := 0.0
		fmt.Println("How much would you like to deposit?")
		fmt.Scan(&amount)

		mu.Lock()
		bb.deposit(amount)
		bb.display()
		mu.Unlock()
	} else if choice == 2 {
		amount := 0.0
		fmt.Println("How much would you like to withdraw?")
		fmt.Scan(&amount)

		mu.Lock()
		bb.withdraw(amount)
		bb.display()
		mu.Unlock()
	} else if choice == 3 {
		mu.Lock()
		bb.display()
		mu.Unlock()
	}
}

func (bb *bankBalance) deposit(amount float64) {
	bb.amount = bb.amount + amount
}

func (bb *bankBalance) withdraw(amount float64) {
	bb.amount = bb.amount - amount
}

func (bb *bankBalance) display() {
	fmt.Printf("You have %s$%.2f remaining in your balance.\n", bb.currency, bb.amount)
}
