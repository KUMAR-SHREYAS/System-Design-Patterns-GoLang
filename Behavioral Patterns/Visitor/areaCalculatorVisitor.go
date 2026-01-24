package main

import (
	"fmt"
	"math"
)

type AreaCalculator struct {
}

func NewAreaCalculator() *AreaCalculator {
	return &AreaCalculator{}
}
func (a *AreaCalculator) visitCircle(circle *Circle) {
	area := math.Pi * circle.getRadius() * circle.getRadius()
	fmt.Printf("Area of Circle is: %f", area)
	fmt.Println()
}

func (a *AreaCalculator) visitRectangle(rect *Rectangle) {
	area := rect.height * rect.width
	fmt.Printf("Area of Rectangle: %f", area)
	fmt.Println()
}
