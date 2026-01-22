// concrete strategy Flat Rate Shipping
package main

import "fmt"

type DistanceBasedShipping struct {
	ratePerKm int
}

func NewDistanceBasedShipping(ratePerKm int) *DistanceBasedShipping {
	return &DistanceBasedShipping{ratePerKm: ratePerKm}
}

func (wbs *DistanceBasedShipping) CalculateCost(order Order) float64 {
	fmt.Println("Calculating with Distance Based strategy for " + order.getDestinationZone())
	var finalRatePerkm float64
	switch order.getDestinationZone() {
	case "Zone A":
		finalRatePerkm = float64(wbs.ratePerKm) * 5.0
	case "Zone B":
		finalRatePerkm = float64(wbs.ratePerKm) * 7.0
	default:
		finalRatePerkm = float64(wbs.ratePerKm) * 10.0
	}
	return finalRatePerkm * order.getOrderWeight()
}
