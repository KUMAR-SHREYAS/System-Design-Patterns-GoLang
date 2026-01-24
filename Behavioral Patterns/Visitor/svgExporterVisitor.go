package main

import "fmt"

type SVGExporter struct {
}

func NewSVGExporter() *SVGExporter {
	return &SVGExporter{}
}

func (s *SVGExporter) visitCircle(circle *Circle) {
	fmt.Printf("<circle r=\"%v\" />\n", circle.getRadius())
}
func (s *SVGExporter) visitRectangle(rect *Rectangle) {
	fmt.Printf("<rect w=\"%v\" h=\"%v\" />\n", rect.getWidth(), rect.getHeight())
}
