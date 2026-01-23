package main

import "fmt"

type IdleState struct {
}

func NewIdleState() *IdleState {
	return &IdleState{}
}

func (i *IdleState) SelectItem(ctx *VendingMachine, itemCode string) {
	fmt.Println("Item Selected: " + itemCode)
	ctx.setSelectedItem(itemCode)
	ctx.setState(NewItemSelectedState())
}

func (i *IdleState) InsertCoin(ctx *VendingMachine, amount float64) {
	ctx.addInsertedAmount(amount)
	fmt.Printf("Inserted $%.2f\n", amount)
}

func (i *IdleState) DispenseItem(ctx *VendingMachine) {
	fmt.Println("No item selected. Nothing to dispense.")
}
