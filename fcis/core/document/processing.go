package document

import (
	"time"
)

// ProcessingResult represents the outcome of document processing
type ProcessingResult struct {
	Document    Document
	ProcessedAt time.Time
}
