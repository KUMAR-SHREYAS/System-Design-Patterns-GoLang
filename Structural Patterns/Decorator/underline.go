package main

import "fmt"

type UnderlineDecorator struct {
	*TextDecorator
}

func NewUnderlineDecorator(inner TextView) *UnderlineDecorator {
	return &UnderlineDecorator{TextDecorator: NewTextDecorator((inner))}
}

func (b *UnderlineDecorator) Render() {
	fmt.Println("<u>")
	b.inner.Render()
	fmt.Println("</u>")
}
