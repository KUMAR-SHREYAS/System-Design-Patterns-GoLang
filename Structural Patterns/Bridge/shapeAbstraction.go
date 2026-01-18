package main

type ShapeMethod interface {
	Draw() //ShapeMethod doesnt care about the type of object, just cares
	// if it implements Draw()
}

type Shape struct {
	renderer Renderer
}

func NewShape(renderer Renderer) *Shape {
	return &Shape{renderer: renderer}
}
