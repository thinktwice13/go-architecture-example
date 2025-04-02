package document

import (
	"time"
)

// ProcessingResult represents the outcome of document processing
type ProcessingResult struct {
	Document      Document
	WordCount     int
	ProcessedAt   time.Time
	ExtractedData map[string]interface{}
}
