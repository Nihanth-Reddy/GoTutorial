package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	fmt.Println("Welcome to Interfaces")
	c := Circle{4}
	r := Rectangle{2, 4}
	fmt.Println("Circle area is: ", c.area(), " and perimeter is: ", c.perimeter())
	fmt.Println("Rectangle area is: ", r.area(), " and perimeter is: ", r.perimeter())

	shapes := []Shape{&c, &r}
	for _, shape := range shapes {
		fmt.Printf("Perimeter of the shape %T is %v\n", shape, shape.perimeter())
	}
}
