// concrete factory adidas
package main

type nike struct {
}

func (a *nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 13,
		},
	}
}

func (a *nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 34,
		},
	}
}
