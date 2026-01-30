package structs_methods_interfaces

import (
	"fmt"
	"testing"
)

func TestShapesArea(t *testing.T) {
	testCases := []struct {
		title        string
		shape        Shape
		expectedArea float64
	}{
		{title: "Square.Area should produce correct result", shape: &Square{Side: 40.0}, expectedArea: 1600.0},
		{title: "Circle.Area should produce correct result", shape: &Circle{Radius: 25.0}, expectedArea: 1963.50},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			result := tc.shape.Area()
			fmt.Println("Result: ", result)
			if result != tc.expectedArea {
				t.Errorf("Expected area: %.2f, got: %.2f", tc.expectedArea, result)
			}
		})
	}
}

func TestShapesPerimeter(t *testing.T) {
	testCases := []struct {
		title         string
		shape         Shape
		expectedPerim float64
	}{
		{title: "Square.Perimeter should produce correct result", shape: &Square{Side: 40.0}, expectedPerim: 160.0},
		{title: "Circle.Perimeter should produce correct result", shape: &Circle{Radius: 25.0}, expectedPerim: 157.08},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			result := tc.shape.Perimeter()
			fmt.Println("Result: ", result)
			if result != tc.expectedPerim {
				t.Errorf("Expected area: %.2f, got: %.2f", tc.expectedPerim, result)
			}
		})
	}
}

// Your standard way of doing testing
// func TestStandard(t *testing.T) {
// 	t.Run("Square.Area should produce correct result", func(t *testing.T) {
// 		square := Square{
// 			Side: 25.0,
// 		}
// 		result := square.Area()
// 		expected := 625.0

// 		if result != expected {
// 			t.Errorf("Got %.2f, expected %.2f", result, expected)
// 		}
// 	})

// 	t.Run("Square.Perimeter should produce correct result", func(t *testing.T) {
// 		square := Square{
// 			Side: 25.0,
// 		}
// 		result := square.Perimeter()
// 		expected := 100.0

// 		if result != expected {
// 			t.Errorf("Got %.2f, expected %.2f", result, expected)
// 		}
// 	})

// 	t.Run("Circle.Area should produce correct result", func(t *testing.T) {
// 		circle := Circle{
// 			Radius: 8.5,
// 		}
// 		result := circle.Area()
// 		expected := 625.0

// 		if result != expected {
// 			t.Errorf("Got %.2f, expected %.2f", result, expected)
// 		}
// 	})

// 	t.Run("Circle.Diameter should produce correct result", func(t *testing.T) {
// 		square := Square{
// 			Side: 25.0,
// 		}
// 		result := square.Perimeter()
// 		expected := 100.0

// 		if result != expected {
// 			t.Errorf("Got %.2f, expected %.2f", result, expected)
// 		}
// 	})

// 	t.Run("Circle.Perimeter should produce correct result", func(t *testing.T) {
// 		square := Square{
// 			Side: 25.0,
// 		}
// 		result := square.Perimeter()
// 		expected := 100.0

// 		if result != expected {
// 			t.Errorf("Got %.2f, expected %.2f", result, expected)
// 		}
// 	})
// }
