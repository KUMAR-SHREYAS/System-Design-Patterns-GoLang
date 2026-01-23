package main

import "fmt"

type ItemSelected struct {
}

func NewItemSelectedState() *ItemSelected {
	return &ItemSelected{}

}
func (i *ItemSelected) SelectItem(ctx *VendingMachine, itemCode string) {
	fmt.Printf("Item already selected: %s", ctx.getSelectedItem())
}

func (i *ItemSelected) InsertCoin(ctx *VendingMachine, amt float64) {
	ctx.addInsertedAmount(amt)
	fmt.Printf("Inserted $%.2f for item: %s, total: $%.2f\n", amt, ctx.getSelectedItem(), ctx.insertedAmount)
	ctx.setState(NewHasMoneyState())
}

func (i *ItemSelected) DispenseItem(ctx *VendingMachine) {
	fmt.Println("Insert coin before dispensing.")
}
