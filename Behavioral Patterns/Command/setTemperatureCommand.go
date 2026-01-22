package main

type SetTemperature struct {
	thermostat Thermostat
	newTemp    int
	prevTemp   int
}

func NewSetTemperature(thermostat Thermostat, temp int) *SetTemperature {
	return &SetTemperature{
		thermostat: thermostat,
		newTemp:    temp,
	}
}

func (s *SetTemperature) Execute() {
	s.prevTemp = s.thermostat.getCurrentTemperature()
	s.thermostat.setTemperature(s.newTemp)
}

func (s *SetTemperature) Undo() {
	s.thermostat.setTemperature(s.prevTemp)

}
