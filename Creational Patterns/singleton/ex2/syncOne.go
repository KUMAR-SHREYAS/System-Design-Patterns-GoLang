package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct { //class body
}

var singleInstance *single // single object creation

func getInstance() *single { //only once initialization
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}
