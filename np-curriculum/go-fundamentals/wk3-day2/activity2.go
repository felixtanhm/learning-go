package main

import (
	"fmt"
	"strconv"
)

func squares(digits []int, result chan int) {
	squares := 0
	for _, v := range digits {
		squares = squares + (v * v)
	}
	result <- squares
}

func cubes(digits []int, result chan int) {
	cubes := 0
	for _, v := range digits {
		cubes = cubes + (v * v * v)
	}
	result <- cubes
}

func splitNum(num int) []int {
	digits := []int{}
	charArr := strconv.Itoa(num)
	for _, char := range charArr {
		num, _ := strconv.Atoi(string(char))
		digits = append(digits, num)
	}
	return digits
}

func main() {
	chnl := make(chan int, 2)
	result := 0
	input := 0

	fmt.Println("Please provide a positive integer")
	fmt.Scan(&input)
	numArr := splitNum(input)

	go squares(numArr, chnl)
	go cubes(numArr, chnl)

	for i := 0; i < 2; i++ {
		result += <-chnl
	}
	fmt.Println("The final result is:", result)
	close(chnl)
}
