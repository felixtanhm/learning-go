package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("Result: %d | Expected: %d", result, expected)
	}
}

func ExampleAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}
