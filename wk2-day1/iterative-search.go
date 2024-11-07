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

	sequentialIndex, sequentialComparisons := searchSorted(numArr, len(numArr), target)
	binaryIndex, binaryComparisons := binarySearch(numArr, len(numArr), target)
	fmt.Println("\nSequential Index of target is: ", sequentialIndex)
	fmt.Println("Number of comparisons required: ", sequentialComparisons)

	fmt.Println("\nBinary Index of target is: ", binaryIndex)
	fmt.Println("Number of comparisons required: ", binaryComparisons)

	fmt.Println("\nSearch Comparison Results:")
	fmt.Println("+---------+----------------------+------------------+")
	fmt.Println("| Target  | Sequential Search    | Binary Search    |")
	fmt.Println("+---------+----------------------+------------------+")
	targets := []int{1, 2, 99, 100, 999, 1000, 1999, 2000}
	for _, t := range targets {
		_, seqComp := searchSorted(numArr[:], len(numArr), t)
		_, binComp := binarySearch(numArr[:], len(numArr), t)
		fmt.Printf("| %-7d | %-20d | %-16d |\n", t, seqComp, binComp)
	}
	fmt.Println("+---------+----------------------+------------------+")
}

func searchSorted(arr []int, arraySize int, target int) (int, int) {
	comparisons := 0
	for i, v := range arr {
		comparisons++
		if v == target {
			return i, comparisons
		}
	}
	return -1, comparisons
}

func binarySearch(arr []int, arraySize int, target int) (int, int) {
	left, right, comparisons := 0, arraySize-1, 0
	for left <= right {
		comparisons++
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid, comparisons
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, comparisons
}
