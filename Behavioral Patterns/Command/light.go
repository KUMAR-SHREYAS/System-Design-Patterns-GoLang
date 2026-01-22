package main

import "fmt"

type Light struct {
}

func NewLight() Light {
	return Light{}
}
func (l *Light) on() {
	fmt.Println("Light turned ON")
}

func (l *Light) off() {
	fmt.Println("Light turned OFF")
}
