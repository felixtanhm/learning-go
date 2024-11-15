package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	area, err := calCircleArea(-1)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Area of circle is: ", area)
	}
}

func calCircleArea(radius float64) (float64, error) {
	if radius <= 0 {
		return 0, errors.New("Radius needs to be a positive integer.")
	}
	area := math.Pi * math.Pow(radius, 2)
	return area, nil
}
