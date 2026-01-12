package main

import "fmt"

type ItalicDecorator struct {
	*TextDecorator
}

func NewItalicDecorator(inner TextView) *ItalicDecorator {
	return &ItalicDecorator{TextDecorator: NewTextDecorator((inner))}
}

func (b *ItalicDecorator) Render() {
	fmt.Println("<i>")
	b.inner.Render()
	fmt.Println("</i>")
}
