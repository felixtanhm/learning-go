package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Empty Hello", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello World"

		if result != expected {
			t.Errorf("Result: %q | Expected: %q", result, expected)
		}
	})

	t.Run("Dynamic Name", func(t *testing.T) {
		result := Hello("Hello", "Felix")
		expected := "Hello Felix"

		if result != expected {
			t.Errorf("Result: %q | Expected: %q", result, expected)
		}
	})

	t.Run("Dynamic Greeting with Dynamic Name", func(t *testing.T) {
		result := Hello("Bye", "Felix")
		expected := "Bye Felix"

		if result != expected {
			t.Errorf("Result: %q | Expected: %q", result, expected)
		}
	})

}
