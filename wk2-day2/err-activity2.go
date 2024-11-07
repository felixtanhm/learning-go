package main

import (
	"fmt"
	"math"
)

type invalidTriangleError struct {
	a, b, c float64
	msg     string
}

type invalidSideError struct {
	side float64
	msg  string
}

func (e *invalidTriangleError) Error() string {
	return fmt.Sprintf("Invalid triangle with sides a=%.2f, b=%.2f, c=%.2f: %s", e.a, e.b, e.c, e.msg)
}

func (e *invalidSideError) Error() string {
	return fmt.Sprintf("Invalid side %.2f: %s", e.side, e.msg)
}

func main() {
	area, err := createTriangle(20, 12, 5)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Area of triangle is: ", area)
	}
}

func createTriangle(a, b, c float64) (float64, error) {
	if a <= 0 {
		return 0, &invalidSideError{a, "Invalid Side a. Needs to be a positive integer."}
	}
	if b <= 0 {
		return 0, &invalidSideError{a, "Invalid Side b. Needs to be a positive integer."}
	}
	if c <= 0 {
		return 0, &invalidSideError{a, "Invalid Side c. Needs to be a positive integer."}
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return 0, &invalidTriangleError{a, b, c, "The sum of any two sides must be greater than the third side"}
	}

	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return area, nil
}
