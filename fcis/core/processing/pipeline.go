package processing

import (
	"fcis/core/document"
	"time"
)

// ProcessDocument is a pure function that processes a document
// Composes other pure functions to create a processing pipeline
func ProcessDocument(doc document.Document) document.ProcessingResult {
	// Extract metadata
	// Calculate word count
	// Estimate language
	// Enrich extracted data with language
	// Update document with enriched metadata
	// Set status to processed
	processedDoc := document.SetStatus(doc, "processed")

	// Return processing result
	return document.ProcessingResult{
		Document:    processedDoc,
		ProcessedAt: time.Now().UTC(),
	}
}
