package shapes

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (circle Circle) Area() float64 {
	return circle.radius * circle.radius * math.Pi
}

func (tri Triangle) Area() float64 {
	return tri.base * tri.height * 0.5
}
