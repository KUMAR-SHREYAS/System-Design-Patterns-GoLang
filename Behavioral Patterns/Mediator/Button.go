package main

import "fmt"

type Button struct {
	component UIComponent
	enabled   bool
}

func NewButton(mediator UIMediator) *Button {
	comp := NewUIComponent(mediator)
	return &Button{component: *comp, enabled: true}
}
func (b *Button) enable() {
	b.setEnabled(true)
}
func (b *Button) setEnabled(val bool) {
	b.enabled = val
	status := "DISABLED"
	if b.enabled {
		status = "ENABLED"
	}
	fmt.Println("Login Button is now ", status)
}
func (b *Button) click() {
	if b.enabled {
		fmt.Println("Login buton clicked!")
		b.component.notifyMediator()
	} else {
		fmt.Println("Login Button is Disabled.")
	}
}
