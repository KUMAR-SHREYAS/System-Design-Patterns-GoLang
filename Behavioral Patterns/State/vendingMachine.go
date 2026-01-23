package main

type VendingMachine struct {
	currState      MachineState
	selectedItem   string
	insertedAmount float64
	itemPrices     map[string]float64
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{
		currState:  NewIdleState(),
		itemPrices: map[string]float64{"A1": 1.0, "B2": 1.5},
	}
}

func (v *VendingMachine) setState(state MachineState) {
	v.currState = state
}
func (v *VendingMachine) setSelectedItem(itemCode string) {
	v.selectedItem = itemCode
}

func (v *VendingMachine) setInsertedAmount(amt float64) {
	v.insertedAmount = amt
}
func (v *VendingMachine) addInsertedAmount(amt float64) {
	v.insertedAmount += amt
}
func (v *VendingMachine) getSelectedItem() string {
	return v.selectedItem
}
func (v *VendingMachine) getItemPrice(itemCode string) float64 {
	return v.itemPrices[itemCode]
}
func (v *VendingMachine) selectItem(itemCode string) {
	v.currState.SelectItem(v, itemCode)
}
func (v *VendingMachine) insertCoin(amt float64) {
	v.currState.InsertCoin(v, amt)
}
func (v *VendingMachine) dispenseItem() {
	v.currState.DispenseItem(v)
}
func (v *VendingMachine) reset() {
	v.selectedItem = ""
	v.insertedAmount = 0.0
	v.currState = NewIdleState()
}
