package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{1.0, 2.0}
	result := Perimeter(rectangle)
	expected := 6.0

	if result != expected {
		t.Errorf("Result: %.2f | Expected: %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{1.0, 2.0}, expected: 2.0},
		{
			name:     "Circle",
			shape:    Circle{10.0},
			expected: 314.1592653589793,
		},
		{
			name:     "Triangle",
			shape:    Triangle{12, 6},
			expected: 36.0,
		},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.shape.Area()
			if result != test.expected {
				t.Errorf("Result %g | Expected %g | Input %#v", result, test.expected, test.shape)
			}
		})
	}

	// checkArea := func(t testing.TB, shape Shape, expected float64) {
	// 	t.Helper()
	// 	result := shape.Area()
	// 	if result != expected {
	// 		t.Errorf("Result %g | Expected %g", result, expected)
	// 	}
	// }

	// t.Run("Rectangle Area", func(t *testing.T) {
	// 	rectangle := Rectangle{1.0, 2.0}
	// 	checkArea(t, rectangle, 2.0)
	// })

	// t.Run("Circle Area", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
