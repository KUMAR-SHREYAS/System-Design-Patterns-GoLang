package main

// TextDecorator wraps TextView Component Interaface
// hence TextDecorator also implements TextView interface
type TextDecorator struct {
	inner TextView // inner component is TextView
}

func NewTextDecorator(inner TextView) *TextDecorator {
	return &TextDecorator{inner: inner}
}
