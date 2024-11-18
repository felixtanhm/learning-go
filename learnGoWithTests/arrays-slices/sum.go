package arraysslices

func Sum(nums []int) (result int) {
	result = 0
	for _, value := range nums {
		result += value
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, values := range numbersToSum {
		sums = append(sums, Sum(values))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, values := range numbersToSum {
		if len(values) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(values[1:]))
		}
	}
	return sums
}
