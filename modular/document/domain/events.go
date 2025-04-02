package domain

import (
	"modular/common/domain"
	"time"
)

var _ domain.DomainEvent = (*DocumentUploaded)(nil)

// DocumentUploaded is an event that occurs when a document is uploaded
type DocumentUploaded struct {
	domain.BaseEvent
	DocumentID int64
}

// EventName returns the name of the event
func (e DocumentUploaded) EventName() string {
	return string(e.Type)
}

// OccurredAt returns the timestamp of the event
func (e DocumentUploaded) OccurredAt() time.Time {
	return e.Timestamp
}
