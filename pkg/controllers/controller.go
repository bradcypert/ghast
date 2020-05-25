package controllers

import (
	"net/http"

	ghastContainer "github.com/bradcypert/ghast/pkg/container"
)

// GhastController should be embedded into consumer controllers
// and provides helper functions for working with the responseWriter, etc.
type GhastController struct{}

// Container returns the DI container associated with the given controller/request pairing.
func (c GhastController) Container(r *http.Request) ghastContainer.Container {
	return r.Context().Value("ghast/container").(ghastContainer.Container)
}

// PathParam Get a Path Parameter from a given request and key
func (c GhastController) PathParam(r *http.Request, key string) interface{} {
	return r.Context().Value(key)
}
