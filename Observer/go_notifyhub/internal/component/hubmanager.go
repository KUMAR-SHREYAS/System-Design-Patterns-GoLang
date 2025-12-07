package component

// Maps from component/HubManager.java
// Implement HubManager that consumes events from EventHub, updates sinks, and notifies observers.

// Expected to depend on service.Graph and model.SinkNode.
import (
	"bus/go_notifyhub/internal/core"
	"bus/go_notifyhub/internal/service"
	"sync"
)

type HubManager struct {
	Graph      *service.Graph
	Publishers map[string]*core.Publisher
	Observers  map[string]*core.Observer
	mu         sync.RWMutex
}

func NewHubManager(g *service.Graph) *HubManager {
	return &HubManager{
		Graph:      g,
		Publishers: make(map[string]*core.Publisher),
		Observers:  make(map[string]*core.Observer),
	}
}

func (hm *HubManager) RegisterPublisher(p *core.Publisher) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.Publishers[p.ID] = p
}

func (hm *HubManager) RegisterObserver(o *core.Observer) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.Observers[o.ID] = o
}

func (hm *HubManager) GetPublisher(id string) *core.Publisher {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	return hm.Publishers[id]
}

// GetRegistry returns a summary for the API
func (hm *HubManager) GetRegistry() map[string]interface{} {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	return map[string]interface{}{
		"publishers": len(hm.Publishers),
		"observers":  len(hm.Observers),
	}
}

// GetG returns the Graph structure (for /hub/graph endpoint)
func (hm *HubManager) GetG() *service.Graph {
	// The Graph has its own internal mutex, so we can return the pointer directly
	return hm.Graph
}

// Get retrieves specific collections based on the type string (for /hub/:type endpoint)
func (hm *HubManager) Get(resourceType string) interface{} {
	hm.mu.RLock()
	defer hm.mu.RUnlock()

	switch resourceType {
	case "publishers":
		// Return a copy or the map itself.
		// Note: Returning the map directly is fine for JSON serialization,
		// but not thread-safe if the caller modifies it.
		// Since Gin just reads it to serialize to JSON, this is acceptable.
		return hm.Publishers
	case "observers":
		return hm.Observers
	default:
		return map[string]string{"error": "Unknown resource type: " + resourceType}
	}
}
