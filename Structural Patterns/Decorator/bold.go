package main

import "fmt"

type BoldDecorator struct {
	*TextDecorator
}

func NewBoldDecorator(inner TextView) *BoldDecorator {
	return &BoldDecorator{TextDecorator: NewTextDecorator((inner))}
}

func (b *BoldDecorator) Render() {
	fmt.Println("<b>")
	b.inner.Render()
	fmt.Println("</b>")
}
