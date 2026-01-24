package main

type ShapeVisitor interface {
	visitCircle(circle *Circle)
	visitRectangle(rectangle *Rectangle)
}
