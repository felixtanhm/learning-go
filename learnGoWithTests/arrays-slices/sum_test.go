package arraysslices

import (
	"fmt"
	"reflect"
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

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2, 3}, []int{3, 3})
	expected := []int{6, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %v | Expected: %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	result := SumAllTails([]int{0, 2}, []int{2, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %v | Expected: %v", result, expected)
	}
}
