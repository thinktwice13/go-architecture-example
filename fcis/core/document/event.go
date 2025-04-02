package document

import "time"

type EventType string

var (
	EventTypeUploaded EventType = "document.uploaded"
)

type Event struct {
	EventType EventType
	Document  Document
	Timestamp time.Time
}
