package main

import "fmt"

type Thermostat struct {
	currTemp int
}

func NewThermostat() Thermostat {
	return Thermostat{}
}
func (t *Thermostat) setTemperature(temp int) {
	fmt.Printf("Thermostat set to %d degree Celsius", temp)
	t.currTemp = temp
}

func (t *Thermostat) getCurrentTemperature() int {
	return t.currTemp
}
