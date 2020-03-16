package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circ := Circle{0, 0, 5}
	rec := Rectangle{10, 5}

	fmt.Printf("Circle Area - %f\n", getArea(circ))
	fmt.Printf("Rectange Area - %f\n", getArea(rec))
}
