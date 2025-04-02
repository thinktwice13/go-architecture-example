package event

// Publisher is an application-wide interface not defined in the domain
type Publisher interface {
	Publish(event DocumentEvent) error
}
