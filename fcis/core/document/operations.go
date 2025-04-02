package document

// ValidationResult represents the outcome of document validation
// Using value types for results rather than errors for pure functional approach
type ValidationResult struct {
	IsValid bool
	Errors  []string
}

func Validate(doc Document) ValidationResult {
	var result ValidationResult

	if doc.Name == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "Document name cannot be empty")
	} else {
		result.IsValid = true
	}

	return result
}

func ValidateForUpload(doc Document) ValidationResult {
	var result ValidationResult

	if doc.Name == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "Document name cannot be empty")
	} else if doc.Status != "new" {
		result.IsValid = false
		result.Errors = append(result.Errors, "Document status must be 'new'")
	} else {
		result.IsValid = true
	}

	return result
}

func ValidateForStorage(doc Document) ValidationResult {
	var result ValidationResult

	if doc.Name == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "Document name cannot be empty")
	} else if len(doc.Content) == 0 {
		result.IsValid = false
		result.Errors = append(result.Errors, "Document content cannot be empty")
	} else if doc.Status != "new" && doc.Status != "uploaded" {
		result.IsValid = false
		result.Errors = append(result.Errors, "Document status must be 'new' or 'uploaded'")
	} else {
		result.IsValid = true
	}

	return result
}
