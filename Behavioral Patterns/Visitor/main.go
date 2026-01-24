package main

import "fmt"

func main() {
	shapes := []Shape{NewCircle(5), NewRectangle(10, 4), NewCircle(2.5)}
	fmt.Println("=== Calculating Areas ===")
	var areaCalculator ShapeVisitor = NewAreaCalculator()
	for _, shape := range shapes {
		shape.Accept(areaCalculator)
	}
	fmt.Println("\n=== Exporting to SVG ===")
	var svgExporter ShapeVisitor = NewSVGExporter()
	for _, shape := range shapes {
		shape.Accept(svgExporter)
	}
}
