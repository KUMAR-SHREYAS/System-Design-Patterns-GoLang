package main

import "fmt"

type Dispensing struct {
}

func NewDispensingState() *Dispensing {
	return &Dispensing{}
}

func (d *Dispensing) SelectItem(context *VendingMachine, itemCode string) {
	fmt.Println("Please Wiat dispensing in progress")
}

func (d *Dispensing) InsertCoin(context *VendingMachine, amt float64) {
	fmt.Println("Please Wiat dispensing in progress")
}
func (d *Dispensing) DispenseItem(ctx *VendingMachine) {
	fmt.Println("Already Dispensing. Please wait. ")
}
