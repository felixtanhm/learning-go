// Activity 4
// Merge Sort can be improved by incorporating Insertion Sort for smaller subarrays. During the recursion, instead of recursively splitting the array all the way down to single-element subarrays, we define a threshold (typically around 10-20 elements). When the subarray size reaches this threshold, we stop splitting and switch to Insertion Sort.
// Insertion Sort is more efficient for small arrays because it has lower overhead, and its time complexity of O(n2) is not an issue for small n. After sorting small subarrays with Insertion Sort, continue with the regular merging process of Merge Sort. The merging step can still be applied efficiently, as Insertion Sort ensures that each small subarray is already sorted.

package main

import "fmt"

func main() {
	numArr := []int{3, 7, 5, 6, 8, 9, 4, 2, 1}
	mergeSort(numArr)
	fmt.Println(numArr)
}

func mergeSort(arr []int) {
	treshold := 2
	arrLength := len(arr)
	if arrLength <= treshold {
		insertionSort(arr)
	} else {
		mid := arrLength / 2
		left := make([]int, mid)
		right := make([]int, arrLength-mid)
		copy(left, arr[:mid])
		copy(right, arr[mid:])

		mergeSort(left)
		mergeSort(right)
		merge(arr, left, right)
	}
}

func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}

}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
