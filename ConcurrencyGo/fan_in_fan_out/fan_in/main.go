package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Start two independent CSV readers.
	// Each call returns a channel that will stream []string rows.
	ch1, err := read("fan_in_fan_out/fan_in/file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v", err))
	}

	ch2, err := read("fan_in_fan_out/fan_in/file2.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file2 %v", err))
	}

	// exit is used to block main from exiting early.
	// If main exits, all goroutines are killed immediately.
	exit := make(chan struct{})

	// merge2 performs FAN-IN:
	// It merges multiple input channels into one output channel.
	chM := merge2(ch1, ch2)

	// Consumer goroutine:
	// Reads from merged channel and prints records.
	go func() {
		for v := range chM {
			fmt.Println(v)
		}

		// Once merged channel is closed, signal main to exit.
		close(exit)
	}()

	// Block main until exit is closed.
	// Prevents premature termination of program.
	<-exit

	fmt.Println("All completed, exiting")
}

// read opens a CSV file and streams its rows over a channel.
// It returns a RECEIVE-ONLY channel.
func read(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %v", err)
	}

	ch := make(chan []string)
	cr := csv.NewReader(f)

	// Start a goroutine to read file asynchronously.
	go func() {
		for {
			record, err := cr.Read()

			// When file ends, close channel and exit goroutine.
			if err == io.EOF {
				close(ch)
				return
			}

			// Send each CSV record to channel.
			ch <- record
		}
	}()

	return ch, nil
}

// merge2 performs FAN-IN:
// It takes multiple input channels and merges them into a single output channel.
func merge2(cs ...<-chan []string) <-chan []string {
	chans := len(cs) // number of input channels

	// wait channel is used to track completion of each sender goroutine.
	// Buffered so senders don’t block when signaling completion.
	wait := make(chan struct{}, chans)

	out := make(chan []string) // merged output channel

	// send forwards data from one input channel to output channel.
	send := func(c <-chan []string) {

		// When this sender goroutine finishes,
		// send a zero-sized signal into wait.
		defer func() { wait <- struct{}{} }()

		// Forward all values from input to output.
		for n := range c {
			out <- n
		}
	}

	// Start one goroutine per input channel.
	for _, c := range cs {
		go send(c)
	}

	// This goroutine waits for all senders to complete.
	go func() {
		for range wait {
			chans--

			// When all senders are done, break loop.
			if chans == 0 {
				break
			}
		}

		// Close output channel only AFTER all senders finish.
		// Prevents "send on closed channel" panic.
		close(out)
	}()

	return out
}