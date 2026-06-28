package main

import (
	"fmt"
	"math"
)

// Factory Method: ShapeFactory creates shapes based on type string.
// New shapes can be added without modifying existing code (Open/Closed Principle).

// Shape is the product interface.
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// Circle implementation
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64  { return 2 * math.Pi * c.Radius }
func (c Circle) String() string      { return fmt.Sprintf("Circle(radius=%.2f)", c.Radius) }

// Rectangle implementation
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64  { return 2 * (r.Width + r.Height) }
func (r Rectangle) String() string      { return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height) }

// ShapeFactory creates shapes by type.
type ShapeFactory struct{}

func (f ShapeFactory) Create(shapeType string, params ...float64) (Shape, error) {
	switch shapeType {
	case "circle":
		if len(params) < 1 {
			return nil, fmt.Errorf("circle requires radius")
		}
		return Circle{Radius: params[0]}, nil
	case "rectangle":
		if len(params) < 2 {
			return nil, fmt.Errorf("rectangle requires width and height")
		}
		return Rectangle{Width: params[0], Height: params[1]}, nil
	default:
		return nil, fmt.Errorf("unknown shape type: %s", shapeType)
	}
}

func main() {
	factory := ShapeFactory{}

	shapes := []struct {
		kind   string
		params []float64
	}{
		{"circle", []float64{5.0}},
		{"rectangle", []float64{3.0, 4.0}},
		{"circle", []float64{2.5}},
	}

	for _, s := range shapes {
		shape, err := factory.Create(s.kind, s.params...)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("%s -> Area: %.2f, Perimeter: %.2f\n", shape, shape.Area(), shape.Perimeter())
	}
}
