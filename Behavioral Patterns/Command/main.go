package main

import "fmt"

func main() {
	light := NewLight()
	thermostat := NewThermostat()

	var lightOn Command = NewLightOnCommand(light)
	var lightOff Command = NewLightOffCommand(light)
	var setTemp22 Command = NewSetTemperature(thermostat, 22)

	button := NewSmartButton()

	fmt.Println("-> Pressing Light ON")
	button.setCommand(lightOn)
	button.press()

	fmt.Println("\n-> Pressing Set Temp t 22 degree celsius")
	button.setCommand(setTemp22)
	button.press()

	fmt.Println("\n-> Pressing Light OFF")
	button.setCommand(lightOff)
	button.press()

	fmt.Println("\n↶ Undo Last Action")
	button.undoLast() // undo Set Temp

	fmt.Println("\n↶ Undo Again")
	button.undoLast() // undo light On

}
