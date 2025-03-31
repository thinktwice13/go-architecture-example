package entity

import "time"

type Document struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

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
