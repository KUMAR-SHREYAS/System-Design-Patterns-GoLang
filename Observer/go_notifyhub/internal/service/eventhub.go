package service

import (
	"sync"

	"bus/go_notifyhub/internal/data"
)

// EventHub provides thread-safe event queuing using Go channels
// Maps from service/EventHub.java (LinkedBlockingQueue<Data> eventBus)
type EventHub struct {
	eventBus chan *data.Data
	mu       sync.Mutex
	closed   bool
}

// NewEventHub creates a new EventHub with a buffered channel
// bufferSize is the capacity (equivalent to LinkedBlockingQueue size)
func NewEventHub(bufferSize int) *EventHub {
	return &EventHub{
		eventBus: make(chan *data.Data, bufferSize),
		closed:   false,
	}
}

// PushIntoHub adds an event to the queue (blocking if full)
// Maps from service/EventHub.java: synchronized boolean pushIntoHub(Data data)
func (eh *EventHub) PushIntoHub(d *data.Data) bool {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	eh.eventBus <- d
	return true
}

// Extract retrieves the next event from the queue (blocking if empty)
// Maps from service/EventHub.java: protected Data extract()
func (eh *EventHub) Extract() (*data.Data, bool) {
	d, ok := <-eh.eventBus
	return d, ok
}

// TryExtract attempts to retrieve an event without blocking
// Non-blocking variant (equivalent to Java's eventBus.poll())
func (eh *EventHub) TryExtract() (*data.Data, bool) {
	select {
	case d := <-eh.eventBus:
		return d, true
	default:
		return nil, false
	}
}

// Close gracefully closes the event hub
func (eh *EventHub) Close() {
	eh.mu.Lock()
	defer eh.mu.Unlock()

	if !eh.closed {
		close(eh.eventBus)
		eh.closed = true
	}
}
