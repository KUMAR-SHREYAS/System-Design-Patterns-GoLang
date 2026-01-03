package core

// Maps from core/Publisher.java
// Define Publisher struct with ID and methods: Subscribe, Unsubscribe, Publish.

// Publisher will hold references to Graph and HubManager to perform actions.

import (
	"bus/go_notifyhub/internal/data"
	"fmt"
	"sync"
)

type Publisher struct {
	ID        string
	observers []ObserverInterface //List of subscribers
	mu        sync.RWMutex
}

func NewPublisher(id string) *Publisher {
	return &Publisher{
		ID:        id,
		observers: make([]ObserverInterface, 0),
	}
}

func (p *Publisher) GetPublisherID() string {
	return p.ID
}

func (p *Publisher) Subscribe(o ObserverInterface) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	p.observers = append(p.observers, o)
	fmt.Printf("[Publisher %s] Observer %s subscribed.\n", p.ID, o.GetObserverID())
}

func (p *Publisher) Notify(d *data.Data) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	fmt.Printf("[Publisher %s] Notifying %d observers...\n", p.ID, len(p.observers))
	for _, o := range p.observers {
		o.Update(d)
	}
}
