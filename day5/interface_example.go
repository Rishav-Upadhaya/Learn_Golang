package day5

import "fmt"

// interface: collection of method signature

type Shape interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	Length  int
	Breadth int
}

// To inherit Shape interface. Rectangle Struct must implemnt all the method inside the interface

func (r Rectangle) Area() int {
	return r.Length * r.Breadth
}
func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Breadth)
}

type Circle struct {
	Radius int
}

func (c Circle) Area() int {
	return 3 * c.Radius * c.Radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.Radius
}

// Polymorphism using Shape interface
func DisplayShapeinfo(s Shape) {
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}
