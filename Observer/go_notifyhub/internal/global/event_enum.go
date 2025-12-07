package global

type EventType string

const (
	EventDataIngest EventType = "DATA_INGEST"
	EventProcess    EventType = "PROCESS"
	EventError      EventType = "ERROR"
)
