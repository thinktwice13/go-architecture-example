package upload

import (
	"hex/core/domain"
	"time"
)

var _ domain.DocumentEvent = (*DocumentUploaded)(nil)

type DocumentUploaded struct {
	DocID     int64
	Timestamp time.Time
}

func (e *DocumentUploaded) GetDocumentID() int64 {
	return e.DocID
}
