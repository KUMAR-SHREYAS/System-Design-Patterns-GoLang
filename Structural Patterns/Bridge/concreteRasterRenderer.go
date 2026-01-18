package main

import "fmt"

type RasterRenderer struct {
}

func NewRasterRenderer() *RasterRenderer {
	return &RasterRenderer{}
}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Println("Drawing a circle of radius: " + fmt.Sprintf("%.2f", radius) + " using RASTER Rendering")
}

func (r *RasterRenderer) RenderRectangle(width, height float64) {
	fmt.Println("Drawing a rectangle of width: " + fmt.Sprintf("%.2f", width) + " and height: " + fmt.Sprintf("%.2f", height) + " using RASTER Rendering")
}
