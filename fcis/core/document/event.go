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

func NewUploadedEvent(doc Document) Event {
	return Event{
		EventType: EventTypeUploaded,
		Document:  doc,
		Timestamp: time.Now().UTC(),
	}
}
