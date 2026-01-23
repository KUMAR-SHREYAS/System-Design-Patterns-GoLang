package main

import "fmt"

type Label struct {
	component UIComponent
	text      string
}

func NewLabel(mediator UIMediator) *Label {
	comp := NewUIComponent(mediator)
	return &Label{component: *comp, text: ""}
}

func (l *Label) setText(message string) {
	l.text = message
	fmt.Printf("Status: %s", message)
}
