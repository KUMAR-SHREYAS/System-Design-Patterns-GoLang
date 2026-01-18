package main

import "fmt"

type VectorRenderer struct {
}

func NewVectorRenderer() *VectorRenderer {
	return &VectorRenderer{}
}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Println("Drawing a circle of radius: " + fmt.Sprintf("%.2f", radius) + " using VECTOR Rendering")
}

func (v *VectorRenderer) RenderRectangle(width, height float64) {
	fmt.Println("Drawing a rectangle of width: " + fmt.Sprintf("%.2f", width) + " and height: " + fmt.Sprintf("%.2f", height) + " using VECTOR Rendering")
}
