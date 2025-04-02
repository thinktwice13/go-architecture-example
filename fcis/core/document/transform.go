package document

// SetStatus is a pure function that returns a new document with updated status
// Functional approach means no mutation, return new values
func SetStatus(doc Document, status string) Document {
	newDoc := doc
	newDoc.Status = status
	return newDoc
}
