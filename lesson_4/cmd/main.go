package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle: radius=%.2f", c.radius)
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle: width=%.2f height=%.2f", r.width, r.height)
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}


func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
func main() {
	c := &Circle{radius: 8}
	r := &Rectangle{
		height: 9,
		width:  3,
	}
	DescribeShape(c)
	DescribeShape(r)

}
