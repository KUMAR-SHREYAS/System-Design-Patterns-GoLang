package main

type ShippingStrategy interface {
	CalculateCost(order Order) float64
}
