package main

import "fmt"

type thirdPartyShipping struct {
	baseFee       float64
	percentageFee float64
}

func NewThirdPartyShipping(baseFee, percentageFee float64) *thirdPartyShipping {
	return &thirdPartyShipping{baseFee: baseFee, percentageFee: percentageFee}
}

func (tps *thirdPartyShipping) CalculateCost(order Order) float64 {
	fmt.Println("Calculating with Third-Party API strategy.")
	return tps.baseFee + (tps.percentageFee * order.getOrderValue())
}
