//client code

package main

import "fmt"

func main() {
	order1 := NewOrder()
	//write some test order data
	order1.weight = 12.0
	order1.destinationZone = "Zone A"
	order1.orderValue = 250.0
	var flatRate ShippingStrategy = NewFlatRateShipping(10.0)
	var weightBased ShippingStrategy = NewWeightBasedShipping(2.5)
	var distanceBased ShippingStrategy = NewDistanceBasedShipping(5.0)
	var thirdParty ShippingStrategy = NewThirdPartyShipping(7.5, 0.02)

	// Create context with an initial strategy
	shippingService := NewShippingCostService(flatRate)

	fmt.Println("--- Order 1: Using Flat Rate (initial) ---")
	shippingService.calculateShippingCost(*order1)

	fmt.Println("\n--- Order 1: Changing to Weight-Based ---")
	shippingService.setStrategy(weightBased)
	shippingService.calculateShippingCost(*order1)

	fmt.Println("\n--- Order 1: Changing to Distance-Based ---")
	shippingService.setStrategy(distanceBased)
	shippingService.calculateShippingCost(*order1)

	fmt.Println("\n--- Order 1: Changing to Third Party Api-Based ---")
	shippingService.setStrategy(thirdParty)
	shippingService.calculateShippingCost(*order1)
}
