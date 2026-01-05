package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{} // create a lock to control access to the instance
type single struct {
}

var singleInstance *single // make a variable to hold the single instance

func getInstance() *single { //return the single instance whenever called
	if singleInstance == nil {
		lock.Lock()                //lock the instance
		defer lock.Unlock()        //unlock the instance after creating the instance , use defer keyword
		if singleInstance == nil { // double check to ensure only one instance is created by multiple goroutines
			fmt.Println("Creating single instance now. ")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}
