package main

import "fmt"

type PlainTextView struct {
	text string
}

func NewPlainTextView(text string) *PlainTextView {
	return &PlainTextView{text: text}
}

func (p *PlainTextView) Render() {
	fmt.Println(p.text)
}
