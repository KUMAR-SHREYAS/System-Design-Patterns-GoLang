package main

import "fmt"

type TextField struct {
	component UIComponent
	text      string
}

func NewTextField(mediator UIMediator) *TextField {
	comp := NewUIComponent(mediator)
	return &TextField{component: *comp, text: ""}
}
func (t *TextField) setText(newText string) {
	t.text = newText
	fmt.Printf("TextField updated: %s", newText)
	fmt.Println()
	t.component.notifyMediator()
}
func (t *TextField) getText() string {
	return t.text
}
