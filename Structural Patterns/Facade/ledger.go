package main

import "fmt"

type Ledger struct{}

func (l *Ledger) makeEntry(accountId string, txnType string, amount int) {
	fmt.Printf("Make Ledger entry fro accountId %s with txnType %s for amount %d\n", accountId, txnType, amount)
	return
}
