package domain

import (
	"time"
)

type EventType string

const (
	EventDocumentUploaded EventType = "document.uploaded"
)

type DocumentEvent struct {
	Type      EventType
	Timestamp time.Time
	Data      []byte
}
