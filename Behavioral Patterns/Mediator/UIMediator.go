package main

type UIMediator interface {
	ComponentChanged(component *UIComponent)
}
