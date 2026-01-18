package main

type Circle struct {
	*Shape
	radius float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{
		Shape:  &Shape{renderer: renderer},
		radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}
