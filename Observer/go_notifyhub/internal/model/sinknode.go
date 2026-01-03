package model

// Maps from model/SinkNode.java
// Implement SinkNode with atomic currentVersion and currentAmount and Update(data *data.Data).

type SinkNode struct {
	ID   string `json:"id"`
	Type string `json:"type"` // e.g. LOG, DB, API
	URL  string `json:"url"`
}
