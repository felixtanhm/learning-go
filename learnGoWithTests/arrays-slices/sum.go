package arraysslices

func Sum(nums []int) (result int) {
	result = 0
	for _, value := range nums {
		result += value
	}
	return result
}
