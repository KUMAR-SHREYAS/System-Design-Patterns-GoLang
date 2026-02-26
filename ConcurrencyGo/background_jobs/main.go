package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	fmt.Println("Process ID", os.Getpid())
	s := NewScheduler(5, 10) // no of workers and the channel buffer size

	s.ListenForWork()

	fmt.Println("ready")
	<-waitToExit() // closes when user enter ctrl + C
	s.Exit()
	fmt.Println("exiting")
}

func waitToExit() <-chan struct{} {
	runC := make(chan struct{}, 1)

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, os.Interrupt)

	go func() {
		defer close(runC)
		<-sc
	}()
	return runC
}

type Scheduler struct {
	workers   int
	msgC      chan struct{}
	signalC   chan os.Signal
	waitGroup sync.WaitGroup
}

func NewScheduler(workers, buffer int) *Scheduler {
	return &Scheduler{
		workers: workers,
		msgC:    make(chan struct{}, buffer),
		signalC: make(chan os.Signal, 1),
	}
}

func (s *Scheduler) ListenForWork() {
	go func() { // 1) Listen for messages to process
		// signal.Notify(s.signalC, syscall.SIGTERM) //SIGTERM behaves unpredictably on windows unike Linux
		for {
			// <-s.signalC
			time.Sleep(2 * time.Second)
			s.msgC <- struct{}{} // 2) Send to processing channel
		}
	}()
	s.waitGroup.Add(s.workers)

	for i := 0; i < s.workers; i++ {
		go func() {
			for {
				select {
				case _, open := <-s.msgC: // 3) Wait for messages to process
					if !open { //closed, exiting
						fmt.Printf("%d closing\n", i+1)
						s.waitGroup.Done()
						return
					}
					fmt.Printf("%d<- Processing\n", i)
				}
			}
		}()
	}
}

func (s *Scheduler) Exit() {
	close(s.msgC)
	s.waitGroup.Wait() // Wait blocks until the WaitGroup task counter is zero.
}
