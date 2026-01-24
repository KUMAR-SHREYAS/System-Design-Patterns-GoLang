package main

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) getWidth() float64 {
	return r.width
}
func (r *Rectangle) getHeight() float64 {
	return r.height
}
func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.visitRectangle(r)
}
