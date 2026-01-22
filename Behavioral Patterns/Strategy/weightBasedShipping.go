// concrete strategy Flat Rate Shipping
package main

import "fmt"

type WeightBasedShipping struct {
	ratePerKg float64
}

func NewWeightBasedShipping(ratePerKg float64) *WeightBasedShipping {
	return &WeightBasedShipping{ratePerKg: ratePerKg}
}

func (wbs *WeightBasedShipping) CalculateCost(order Order) float64 {
	fmt.Printf("Calculating with Weight Based strategy ($%.2f)", float64(wbs.ratePerKg))
	return float64(wbs.ratePerKg) * order.getOrderWeight()
}
