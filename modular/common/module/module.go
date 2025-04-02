package module

import "github.com/julienschmidt/httprouter"

type Interface interface {
	Name() string
	SetupRoutes(*httprouter.Router)
	RegisterEventHandlers()
	Start() error
	Stop() error
}
