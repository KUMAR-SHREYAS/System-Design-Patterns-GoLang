// concrete factory adidas
package main

type adidas struct {
}

func (a *adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 11,
		},
	}
}

func (a *adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 30,
		},
	}
}
