package main

type Order struct {
	weight          float64
	destinationZone string
	orderValue      float64
}

func NewOrder() *Order {
	return &Order{}
}

func (o Order) getOrderWeight() float64 {
	return o.weight
}

func (o Order) getDestinationZone() string {
	return o.destinationZone
}

func (o Order) getOrderValue() float64 {
	return o.orderValue
}
