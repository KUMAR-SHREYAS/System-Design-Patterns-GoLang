package main

type UIComponent struct {
	mediator UIMediator
}

func NewUIComponent(mediator UIMediator) *UIComponent {
	return &UIComponent{mediator: mediator}
}
func (u *UIComponent) notifyMediator() {
	u.mediator.ComponentChanged(u)
}
