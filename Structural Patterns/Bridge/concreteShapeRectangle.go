package main

type Rectangle struct {
	renderer Renderer
	width    float64
	height   float64
}

func NewRectangle(renderer Renderer, width, height float64) *Rectangle {
	return &Rectangle{renderer: renderer, width: width, height: height}
}

func (r *Rectangle) Draw() {
	r.renderer.RenderRectangle(r.width, r.height)
}
