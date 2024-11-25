package di

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(os.Stdout, "Chris")
	Greet(&buffer, "Chris")

	result := buffer.String()
	expected := "Hello, Chris\n"

	if result != expected {
		t.Errorf("Result %s | Expected %s", result, expected)
	}
}
