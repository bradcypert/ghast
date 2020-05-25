package app

import (
	"log"
	"net/http"

	ghastContainer "github.com/bradcypert/ghast/pkg/container"
	ghastRouter "github.com/bradcypert/ghast/pkg/router"
)

// App defines a struct that encapsulates the entire ghast framework + application specific settings
type App struct {
	c            ghastContainer.Container
	serverConfig *http.Server
}

// NewApp constructor function for ghast app
func NewApp() App {
	return App{
		ghastContainer.NewContainer(),
		nil,
	}
}

// Start boots up the HTTP server and binds a route listener
func (a App) Start() {
	router, routerOK := a.c.Make("ghast/router").(ghastRouter.Router)
	if routerOK != true {
		log.Panic("Router was not bound to DI container. Did you call SetRouter on your app?")
	}

	// Use custom HTTP server config if provided
	var s *http.Server
	if a.serverConfig == nil {
		s = router.DefaultServer()
	} else {
		s = a.serverConfig
	}

	// but always overwrite the handler to use the ghast router
	s.Handler = router

	// add in our DI container for the router to have access to
	router.SetDIContainer(a.c)

	log.Fatal(s.ListenAndServe())
}

// SetServerConfig provides the app with a custom use Server configuration
func (a App) SetServerConfig(config *http.Server) {
	a.serverConfig = config
}

// SetRouter sets a user configured ghast router to be used as the application's default router.
func (a App) SetRouter(router ghastRouter.Router) {
	a.c.Bind("ghast/router", func(c ghastContainer.Container) interface{} {
		return router
	})
}
