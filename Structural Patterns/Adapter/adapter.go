package main

import "fmt"

type WindowsAdapter struct {
	windows *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	// fmt.Println("Adapter Service on ")
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windows.InsertIntoUSBPort()
}
