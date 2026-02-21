package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)     // BUffered Channel
	exit := make(chan struct{}) //to signal main() to exit

	// Producer go routine
	go func() {

		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")

			time.Sleep(1 * time.Second)
		}
		fmt.Println(time.Now(), "all completed, leaving")
		close(ch)
	}()

	//Consumer Go routine
	go func() {
		//Uses infinte loop to receive values via channel
		// overcomplicated, use select for multiple channels
		for {
			select {
			case v, open := <-ch:
				if !open {
					close(exit)
					return
				}
				fmt.Println(time.Now(), "received", v)
			}
		}

		//Simplifed receival using range keyword
		// for v := range ch {
		// 	fmt.Println(time.Now(), "received", v)
		// }
	}()

	fmt.Println(time.Now(), "waiting for everything to complete")
	<-exit
	fmt.Println(time.Now(), "exiting")
}
