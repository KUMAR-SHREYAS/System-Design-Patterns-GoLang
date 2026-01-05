package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun IGun) {
	fmt.Printf("Gun: %s", gun.getName())
	fmt.Println()
	fmt.Printf("Power: %d", gun.getPower())
	fmt.Println()
}
