package event

import "time"

type DocumentEvent interface {
	GetDocumentID() int64
}

type DocumentUploaded struct {
	DocumentID int64
	Timestamp  time.Time
}

func (e DocumentUploaded) GetDocumentID() int64 {
	return e.DocumentID
}

type DocumentRetrieved struct {
	DocumentID int64
	Timestamp  time.Time
}

func (e DocumentRetrieved) GetDocumentID() int64 {
	return e.DocumentID
}
