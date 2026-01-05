package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 30; i++ {
		// fmt.Println("Goroutine ", i)
		go getInstance()
	}
	fmt.Scanln()
}
