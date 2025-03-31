package domain

import "time"

type DocumentEvent interface {
	GetDocumentID() int64
}

var _ DocumentEvent = (*DocumentUploaded)(nil)

type DocumentUploaded struct {
	Document  Document
	Timestamp time.Time
}

func (d DocumentUploaded) GetDocumentID() int64 {
	return d.Document.ID
}
