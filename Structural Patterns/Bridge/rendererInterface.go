package main

type Renderer interface {
	RenderCircle(radius float64)
	RenderRectangle(width, height float64)
}
