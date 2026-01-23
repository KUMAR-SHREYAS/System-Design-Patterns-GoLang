package main

import "fmt"

func main() {
	vm := NewVendingMachine()
	vm.insertCoin(1.0)
	vm.selectItem("A1")
	vm.insertCoin(1.5)
	vm.dispenseItem()

	fmt.Println("Second Transaction")
	vm.selectItem("B2")
	vm.insertCoin(2.0)
	vm.dispenseItem()
}
