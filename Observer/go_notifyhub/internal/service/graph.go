package service

// Maps from service/Graph.java
// Implement a concurrent-safe mapping from publisherID -> observers.

// Example API:
// func NewGraph() *Graph
// func (g *Graph) CreateVector(pubID string)
// func (g *Graph) InsertObserver(pubID string, ob *core.Observer)
// func (g *Graph) RemoveObserver(pubID, obsID string)
// func (g *Graph) ActiveObservers(pubID string) []*core.Observer

import (
	"fmt"
	"sync"
)

// Graph manages the connections between the nodes
type Graph struct {
	mu    sync.RWMutex
	nodes map[string][]string // Adjacency list: NodeID -> [ConnectedNodeIDs]
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.nodes[from] = append(g.nodes[from], to)
}

func (g *Graph) GetDownstream(id string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.nodes[id]
}

func (g *Graph) String() string {
	g.mu.RLock()
	defer g.mu.Unlock()
	return fmt.Sprintf("%v", g.nodes)
}
