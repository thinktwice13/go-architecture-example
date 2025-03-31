package domain

import (
	"time"
	"vert/shared/domain"
)

type DocumentEvent interface {
	GetDocumentID() int64
}

var _ domain.DocumentEvent = (*DocumentUploaded)(nil)

type DocumentUploaded struct {
	Document  Document
	Timestamp time.Time
}

func (e DocumentUploaded) GetDocumentID() int64 {
	return e.Document.ID
}
