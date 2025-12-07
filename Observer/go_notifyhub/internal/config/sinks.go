package config

// Maps from config/Sinks.java
// Initialize and return the sink registry: map[data.Event]*model.SinkNode

// Example:
// func NewSinkRegistry() map[data.Event]*model.SinkNode { ... }

import "bus/go_notifyhub/internal/model"

// GetDefaultSinks returns the initial static configuration of sinks
func GetDefaultSinks() []model.SinkNode {
	return []model.SinkNode{
		{ID: "sink-1", Type: "DB", URL: "jdbc:mysql://localhost:3306/mydb"},
		{ID: "sink-2", Type: "LOG", URL: "/var/log/notify_hub.log"},
		{ID: "sink-3", Type: "API", URL: "https://webhook.site/test"},
	}
}
