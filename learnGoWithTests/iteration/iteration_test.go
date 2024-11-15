package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaa"

	if result != expected {
		t.Errorf("Result: %q | Expected: %q", result, expected)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
