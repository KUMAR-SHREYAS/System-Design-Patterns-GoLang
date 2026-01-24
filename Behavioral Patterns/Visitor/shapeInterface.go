package main

type Shape interface {
	Accept(visitor ShapeVisitor)
}
