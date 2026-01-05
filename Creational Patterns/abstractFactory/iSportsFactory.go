// abstract factiry <<interface>>
package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func getSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "nike":
		return &nike{}, nil
	case "adidas":
		return &adidas{}, nil
	default:
		return nil, fmt.Errorf("Wrong brand type passed")
	}
}
