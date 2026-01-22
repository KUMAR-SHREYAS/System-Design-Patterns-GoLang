// concrete strategy Flat Rate Shipping
package main

import "fmt"

type FlatRateShipping struct {
	flatRate float64
}

func NewFlatRateShipping(rate float64) *FlatRateShipping {
	return &FlatRateShipping{flatRate: rate}
}

func (frs *FlatRateShipping) CalculateCost(order Order) float64 {
	fmt.Printf("Calculating with Flat Rate strategy ($%.2f)", frs.flatRate)
	return float64(frs.flatRate)
}
