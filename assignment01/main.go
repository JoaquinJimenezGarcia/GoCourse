package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	calculateArea() float64
}

func main() {
	t := triangle{
		height: 2.0,
		base:   5.0,
	}

	s := square{
		sideLength: 7.0,
	}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.calculateArea())
}

func (t triangle) calculateArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) calculateArea() float64 {
	return s.sideLength * s.sideLength
}
