package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius=%.2f", c.radius)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle: width=%.2f height=%.2f", r.width, r.height)
}

func (c Circle) Area() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("Negative radius\n")
	} else {
		return math.Pi * c.radius * c.radius, nil
	}
}

func (c Circle) Perimeter() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("Negative radius\n")
	} else {
		return 2 * math.Pi * c.radius, nil
	}

}

func (r Rectangle) Area() (float64, error) {
	if r.width < 0 {
		return 0, errors.New("Negative width\n")
	} else if r.height < 0 {
		return 0, errors.New("Negative height\n")
	} else {
			return r.width * r.height, nil
	}
}

func (r Rectangle) Perimeter() (float64, error) {
	if r.width < 0 {
		return 0, errors.New("Negative width\n")
	} else if r.height < 0 {
		return 0, errors.New("Negative height\n")
	} else {
		return (r.width + r.height) * 2, nil
	}
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	resArea, _ := s.Area()
	fmt.Printf("Area: %.2f\n", resArea)
	resPerimeter, _ := s.Perimeter()
	fmt.Printf("Perimeter: %.2f\n", resPerimeter)
}
func main() {
	c := Circle{radius: 8}
	r := Rectangle{
		height: 9,
		width:  3,
	}
	DescribeShape(c)
	DescribeShape(r)

}
