package main

import "fmt"

type HasMoney struct {
}

func NewHasMoneyState() *HasMoney {
	return &HasMoney{}
}

func (h *HasMoney) SelectItem(ctx *VendingMachine, itemCode string) {
	fmt.Println("Cannot change item after inserting money.")
}

func (h *HasMoney) InsertCoin(ctx *VendingMachine, amt float64) {
	ctx.addInsertedAmount(amt)
	fmt.Printf("Added $%.2f, total inserted: $%.2f\n", amt, ctx.insertedAmount)
}

func (h *HasMoney) DispenseItem(ctx *VendingMachine) {
	price := ctx.getItemPrice(ctx.getSelectedItem())
	if ctx.insertedAmount >= price {
		fmt.Printf("Dispensing Item: %s", ctx.getSelectedItem())
		ctx.setState(NewDispensingState())

		change := ctx.insertedAmount - price
		if change > 0 {
			fmt.Printf("Returning change: $%.2f\n", change)
		}

		fmt.Println("Item Dispensed successfully")
		ctx.reset()
	} else {
		fmt.Printf("Insufficient funds. Inserted: $%.2f, Required: $%.2f\n", ctx.insertedAmount, price)
	}
}
