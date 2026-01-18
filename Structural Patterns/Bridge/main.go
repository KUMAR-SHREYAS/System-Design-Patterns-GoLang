package main

func main() {
	vector := NewVectorRenderer()
	raster := NewRasterRenderer()
	// Store circle, rectangle in a common interface
	// it is possible if they implement the method Draw()
	var shapes []ShapeMethod

	shapes = append(shapes, NewCircle(vector, 5))
	shapes = append(shapes, NewCircle(raster, 5))
	shapes = append(shapes, NewRectangle(vector, 10, 4))
	shapes = append(shapes, NewRectangle(raster, 10, 4))
	for _, shape := range shapes {
		shape.Draw()
	}
}
