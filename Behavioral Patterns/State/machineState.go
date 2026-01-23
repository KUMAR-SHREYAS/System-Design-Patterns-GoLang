package main

type MachineState interface {
	SelectItem(context *VendingMachine, itemCode string)
	InsertCoin(context *VendingMachine, amount float64)
	DispenseItem(context *VendingMachine)
}
