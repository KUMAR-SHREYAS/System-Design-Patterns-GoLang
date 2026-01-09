package main

import "fmt"

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	// fmt.Println("Adapter converts Lightning signal to USB-C signal.")
	fmt.Println("USB connector is plugged into windows machine.")
}
