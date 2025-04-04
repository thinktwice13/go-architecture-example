package module

type Interface interface {
	Name() string
	Start() error

	// Sample interactions
	// SetupRoutes(*httprouter.Router)
	// RegisterEventHandlers()
	// Shutdown() error
}
