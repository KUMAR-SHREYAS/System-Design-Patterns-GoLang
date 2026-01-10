package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	// err = walletFacade.deductMoneyFromWallet("abc", 1234, 15)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
