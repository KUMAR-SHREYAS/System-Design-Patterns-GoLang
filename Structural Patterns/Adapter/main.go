package main

import "fmt"

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningPortIntoComputer(mac)
	fmt.Println()
	fmt.Println()
	windowsMachine := &Windows{}
	WindowsAdapter := &WindowsAdapter{
		windows: windowsMachine,
	}
	// WindowsAdapter.InsertIntoLightningPort()
	// use windwos adapter to insert lightning port into windows machine as it implemens
	// Computer interface
	client.InsertLightningPortIntoComputer(WindowsAdapter)

}
