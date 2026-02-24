package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Choose which concurrency approach to use
	// wait := waitGroups()
	wait := errGroup()

	// Block main until the returned channel is closed
	<-wait
}

///////////////////////////////////////////////////////////////
// 1️⃣ Using sync.WaitGroup
///////////////////////////////////////////////////////////////

func waitGroups() <-chan struct{} {
	// This channel is used only as a completion signal.
	// Nothing is sent on it — it is just closed when all work is done.
	ch := make(chan struct{}, 1)

	var wg sync.WaitGroup

	// File paths to read concurrently
	f1 := "D:/System Design/ConcurrencyGo/errGroup/file1.csv"
	f2 := "D:/System Design/ConcurrencyGo/errGroup/file2.csv"
	f3 := "D:/System Design/ConcurrencyGo/errGroup/file3.csv"

	// Loop over files and launch one goroutine per file (fan-out pattern)
	for _, file := range []string{f1, f2, f3} {

		// Tell WaitGroup that one goroutine is starting
		wg.Add(1)

		go func() {
			// Signal that this goroutine is finished when it returns
			defer wg.Done()

			// Read returns a channel that streams CSV records
			ch, err := read(file)
			if err != nil {
				// Errors are only printed — other goroutines continue
				fmt.Printf("error reading %v\n", err)
				return
			}

			// Consume records from the file-specific channel
			for line := range ch {
				fmt.Println(line)
			}
		}()
	}

	// Separate goroutine waits for all workers to finish
	go func() {
		wg.Wait() // Block until all wg.Done() calls happen
		close(ch) // Signal completion to main()
	}()

	return ch
}

///////////////////////////////////////////////////////////////
// 2️⃣ Using errgroup (error-aware concurrency)
///////////////////////////////////////////////////////////////

func errGroup() <-chan struct{} {
	// Channel used only as a completion signal
	ch := make(chan struct{}, 1)

	// errgroup collects errors from goroutines
	var g errgroup.Group

	f1 := "D:/System Design/ConcurrencyGo/errGroup/file1.csv"
	f2 := "D:/System Design/ConcurrencyGo/errGroup/file2.csv"
	f3 := "D:/System Design/ConcurrencyGo/errGroup/file3.csv"

	// Launch one goroutine per file
	for _, file := range []string{f1, f2, f3} {

		g.Go(func() error {
			// read returns a stream (channel) of CSV records
			ch, err := read(file)
			if err != nil {
				// Returning error makes g.Wait() return that error
				return fmt.Errorf("error reading %w", err)
			}

			// Consume and print records
			for line := range ch {
				fmt.Println(line)
			}

			return nil
		})
	}

	// Wait for all goroutines to finish
	go func() {
		// If any goroutine returned an error, it appears here
		if err := g.Wait(); err != nil {
			fmt.Printf("Error reading files %v", err)
		}

		// Signal completion
		close(ch)
	}()

	return ch
}

///////////////////////////////////////////////////////////////
// 3️⃣ read() - Streams CSV records through a channel
///////////////////////////////////////////////////////////////

func read(file string) (<-chan []string, error) {

	// Open file
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}

	// Channel used to stream CSV records
	ch := make(chan []string)

	// Launch goroutine so reading is non-blocking
	go func() {
		cr := csv.NewReader(f)

		for {
			record, err := cr.Read()

			// When EOF is reached, close channel and stop
			if errors.Is(err, io.EOF) {
				close(ch)
				return
			}

			// Send each record to the consumer
			ch <- record
		}
	}()

	// Return read-only channel to caller
	return ch, nil
}
