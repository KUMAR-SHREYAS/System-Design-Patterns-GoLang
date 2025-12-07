package core

// Maps from core/Observer.java
// Define Observer struct, version/amount tracking, and update/sync methods.

// Example:
// type Observer struct { ID string /*...*/ }
import (
	"bus/go_notifyhub/internal/data"
	"fmt"
)

// ObserverInterface defines what an observer must do
type ObserverInterface interface {
	Update(data *data.Data)
	GetObserverID() string
}

type Observer struct {
	ID string
}

func NewObserver(id string) *Observer {
	return &Observer{ID: id}
}

func (o *Observer) GetObserverID() string {
	return o.ID
}

// Update is called when the Publisher notifies its observers
func (o *Observer) Update(d *data.Data) {
	fmt.Printf("[Observer %s] Received data from %s: %v\n", o.ID, d.Source, d.Payload)
}
