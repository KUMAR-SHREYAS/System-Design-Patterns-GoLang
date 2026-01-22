package main

import (
	"fmt"
	"reflect"
)

type ShippingCostService struct {
	strategy ShippingStrategy
}

func NewShippingCostService(strategy ShippingStrategy) *ShippingCostService {
	return &ShippingCostService{strategy: strategy}
}

func (s *ShippingCostService) setStrategy(strategy ShippingStrategy) {
	fmt.Println("ShippingCostService: Strategy changed to " + reflect.TypeOf(strategy).Elem().Name())
	s.strategy = strategy
}
func (s *ShippingCostService) calculateShippingCost(order Order) float64 {
	if s.strategy == nil {
		fmt.Println("Shipping strategy not set.")
		return 0.0
	}
	cost := s.strategy.CalculateCost(order)
	fmt.Printf("ShippingCostService: Final Calculated Shipping Cost: $%.2f (using %s)\n", cost, reflect.TypeOf(s.strategy).Elem().Name())
	return cost
}
