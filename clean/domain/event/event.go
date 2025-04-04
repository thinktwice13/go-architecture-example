package event

import (
	"time"
)

type DocumentEvent interface {
	GetDocumentID() string
}

type DocumentUploaded struct {
	DocumentID string
	Timestamp  time.Time
}

func (e DocumentUploaded) GetDocumentID() string {
	return e.DocumentID
}

type DocumentRetrieved struct {
	DocumentID string
	Timestamp  time.Time
}

func (e DocumentRetrieved) GetDocumentID() string {
	return e.DocumentID
}
