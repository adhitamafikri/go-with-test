package structs_methods_interfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return math.Round(s.Side*s.Side*100) / 100
}

func (s *Square) Perimeter() float64 {
	return math.Round(s.Side*4*100) / 100
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Round(math.Pi*c.Radius*c.Radius*100) / 100
}

func (c *Circle) Diameter() float64 {
	return math.Round(2*c.Radius*100) / 100
}

func (c *Circle) Perimeter() float64 {
	return math.Round(math.Pi*c.Diameter()*100) / 100
}
