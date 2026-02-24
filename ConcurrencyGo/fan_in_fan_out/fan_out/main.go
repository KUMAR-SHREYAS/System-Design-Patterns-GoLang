package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Read CSV file and get a read-only channel of records
	ch1, err := read("D:/System Design/ConcurrencyGo/fan_in_fan_out/fan_out/file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v", err))
	}

	// Start 3 worker goroutines consuming from SAME input channel (fan-out pattern)
	// Each worker returns a channel that signals when it finishes.
	br1 := split("1", ch1)
	br2 := split("2", ch1)
	br3 := split("3", ch1)

	// Wait until ALL workers finish
	for {
		// If all worker channels are nil, all workers are done
		if br1 == nil && br2 == nil && br3 == nil {
			break
		}

		// Select waits on whichever worker finishes first
		select {

		// Try receiving from worker 1 completion channel
		case _, ok := <-br1:
			// If channel is closed, worker is done
			if !ok {
				br1 = nil // Set to nil so select ignores it
			}

		// Try receiving from worker 2 completion channel
		case _, ok := <-br2:
			if !ok {
				br2 = nil
			}

		// Try receiving from worker 3 completion channel
		case _, ok := <-br3:
			if !ok {
				br3 = nil
			}
		}
	}

	fmt.Println("All completed, exiting")
}

// split represents a worker.
// It reads from the shared input channel and processes records.
// Returns a channel that signals completion.
func split(worker string, ch <-chan []string) chan struct{} {

	// Channel used only to signal completion (no data sent)
	chE := make(chan struct{})

	go func() {
		// Range automatically stops when input channel is closed
		for v := range ch {
			// Simulate processing
			fmt.Println(worker, v)
		}

		// When input channel closes and loop ends,
		// signal that this worker is done
		close(chE)
	}()

	return chE
}

// read reads a CSV file and streams each record into a channel.
// It returns a read-only channel of []string.
func read(file string) (<-chan []string, error) {

	// Open file
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %v", err)
	}

	// Channel to stream CSV rows
	ch := make(chan []string)

	cr := csv.NewReader(f)

	// Launch goroutine to read file asynchronously
	go func() {
		for {
			record, err := cr.Read()

			// When file ends, close channel to signal no more data
			if err == io.EOF {
				close(ch)
				return
			}

			// Send record to workers
			ch <- record
		}
	}()

	return ch, nil
}