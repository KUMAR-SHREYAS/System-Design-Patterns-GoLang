package main

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) getRadius() float64 {
	return c.radius
}
func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.visitCircle(c)
}
