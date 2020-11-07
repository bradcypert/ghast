package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet"
	"github.com/bradcypert/ghast/pkg/config"
	ghastContainer "github.com/bradcypert/ghast/pkg/container"
	ghastRouter "github.com/bradcypert/ghast/pkg/router"
)

// AppContext to be used by Globally required application objects.
// Warning: this context is essential to Ghast's ability to function correctly.
// Override values in this context at your own risk.
var AppContext context.Context

// App defines a struct that encapsulates the entire ghast framework + application specific settings
type App struct {
	c            *ghastContainer.Container
	serverConfig *http.Server
	views        *jet.Set
}

// NewApp constructor function for ghast app
func NewApp() App {
	var root, _ = os.Getwd()
	var views = jet.NewHTMLSet(filepath.Join(root, "views"))
	container := ghastContainer.NewContainer()

	// Bind the config options into the app. This structure can be any number of items deep.
	fmt.Printf("Reading config from %s/config.yml", root)
	configOptions, err := config.Parse(root + "/config.yml")
	if err != nil {
		log.Panic("Unable to bind your yaml config into the Ghast Container. Please ensure that your config is valid YAML")
	}
	configs, err := config.ParsedConfigToContainerKeys(configOptions)
	if err != nil {
		log.Panic("Unable to bind your yaml config into the Ghast Container. Please ensure that your config is valid YAML")
	}

	for k, v := range configs {
		container.Bind("@"+k, func(c *ghastContainer.Container) interface{} {
			return v
		})
	}

	AppContext = context.WithValue(context.Background(), "ghast/container", container)
	return App{
		container,
		nil,
		views,
	}
}

// GetApp gets the app instance out of a given container
func GetApp(c *ghastContainer.Container) App {
	return c.Make("ghast/app").(App)
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
	s.Addr = a.c.Make("@ghast.config.port").(string)

	// Bind the app to the container so its available
	a.c.Bind("ghast/app", func(c *ghastContainer.Container) interface{} {
		return a
	})

	// add in our DI container for the router to have access to
	router.SetDIContainer(a.c)

	log.Fatal(s.ListenAndServe())
}

// SetServerConfig provides the app with a custom use Server configuration
func (a App) SetServerConfig(config *http.Server) {
	a.serverConfig = config
}

// GetViewSet Gets the Application's JET view set
func (a App) GetViewSet() *jet.Set {
	return a.views
}

// SetRouter sets a user configured ghast router to be used as the application's default router.
func (a App) SetRouter(router ghastRouter.Router) {
	a.c.Bind("ghast/router", func(c *ghastContainer.Container) interface{} {
		return router
	})
}
