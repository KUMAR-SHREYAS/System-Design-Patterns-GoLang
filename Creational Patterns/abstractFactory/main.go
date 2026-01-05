package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	fmt.Println("Print Shoe Details")
	printShoeDetails(nikeShoe)
	// fmt.Println("Print Shoe Details")
	printShoeDetails(adidasShoe)

	fmt.Println("Print Shirt Details")
	printShirtDetails(nikeShirt)
	// fmt.Println("Print Shirt Details")
	printShirtDetails(adidasShirt)

}

func printShirtDetails(brand IShirt) {
	fmt.Printf("logo: %s", brand.getLogo())
	fmt.Println()
	fmt.Printf("size: %d", brand.getSize())
	fmt.Println()
}

func printShoeDetails(brand IShoe) {
	fmt.Printf("logo: %s", brand.getLogo())
	fmt.Println()
	fmt.Printf("size: %d", brand.getSize())
	fmt.Println()
}
