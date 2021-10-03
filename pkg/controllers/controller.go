package controllers

import (
	"net/http"

	"github.com/CloudyKit/jet"
	ghastApp "github.com/bradcypert/ghast/pkg/app"
	ghastContainer "github.com/bradcypert/ghast/pkg/container"
)

// GhastController should be embedded into consumer controllers
// and provides helper functions for working with the responseWriter, etc.
type GhastController struct{}

// Container returns the DI container associated with the given controller/request pairing.
func (c GhastController) Container() *ghastContainer.Container {
	return ghastApp.AppContext.Value("ghast/container").(*ghastContainer.Container)
}

// Config gets a config value from the controller's container.
// Config keys map to YAML in a flattened dot structure prefixed by an @.
// For example:
// a:
//   b: "c"
// "c" can be retrieved via @a.b
// We can't guarantee the type, so we return interface{}
func (c GhastController) Config(key string) interface{} {
	return c.Container().Make(key)
}

// PathParam Get a Path Parameter from a given request and key
func (c GhastController) PathParam(r *http.Request, key string) interface{} {
	return r.Context().Value(key)
}

// PathParam Get a Path Parameter from a given request and key
// returns a list of strings as it supports multiple values for a given path param
func (c GhastController) QueryParam(r *http.Request, key string) []string {
	return r.URL.Query()[key]
}

// View executes a view from the app templates
func (c GhastController) View(name string, w http.ResponseWriter, vars jet.VarMap, contextualData interface{}) {
	tmpl, err := ghastApp.GetApp(c.Container()).GetViewSet().GetTemplate(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err = tmpl.Execute(w, vars, contextualData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
