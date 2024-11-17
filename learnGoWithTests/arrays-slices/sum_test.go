package arraysslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Dynamic slice", func(t *testing.T) {
		nums := []int{1, 2, 3}

		result := Sum(nums)
		expected := 6

		if result != expected {
			t.Errorf("Result %d | Expected: %d | Input: %v", result, expected, nums)
		}
	})
}

func ExampleSum() {
	nums := []int{1, 2, 3, 4, 5}
	result := Sum(nums)
	fmt.Println(result)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(nums)
	}
}
