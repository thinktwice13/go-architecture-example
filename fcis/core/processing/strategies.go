package processing

// ProcessingStrategy defines different processing approaches
type ProcessingStrategy func([]byte) map[string]any

// TextProcessingStrategy processes text documents
func TextProcessingStrategy(content []byte) map[string]any {
	result := make(map[string]any)

	// Extract text-specific metadata ...
	result["format"] = "text"
	result["character_count"] = len(content)

	return result
}

// PDFProcessingStrategy processes PDF documents
func PDFProcessingStrategy(content []byte) map[string]any {
	result := make(map[string]any)

	// Extract PDF-specific metadata ...
	// In a real app, this would parse PDF structure
	result["format"] = "pdf"

	return result
}

// GetProcessingStrategy selects the appropriate strategy based on content
func GetProcessingStrategy(contentType string) ProcessingStrategy {
	switch contentType {
	case "application/pdf":
		return PDFProcessingStrategy
	default:
		return TextProcessingStrategy
	}
}
