package main

import "fmt"

func main() {
	numArr := []int{}
	for i := 1; i < 1000; i++ {
		numArr = append(numArr, i*2)
	}

	target := 0
	fmt.Println("Enter an even number between 2 and 2000 to search")
	_, err := fmt.Scan(&target)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer.")
		return
	}

	sequentialIndex := searchSortedR(numArr, 0, 0, target)
	binaryIndex := binarySearchR(numArr, 0, len(numArr), target)
	fmt.Println("\nSequential Index of target is: ", sequentialIndex)

	fmt.Println("Binary Index of target is: ", binaryIndex)
}

func searchSortedR(arr []int, n int, start int, target int) int {
	if n >= len(arr) {
		return -1
	}
	if arr[n] == target {
		return n
	}
	return searchSortedR(arr, n+1, start, target)
}

func binarySearchR(arr []int, first int, last int, target int) int {
	if first > last {
		return -1
	}
	mid := first + (last-first)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchR(arr, mid+1, last, target)
	} else {
		return binarySearchR(arr, first, mid-1, target)
	}
}
