package controllers

import (
	"fmt"
	"net/http"
)

// GhastController should be embedded into consumer controllers
// and provides helper functions for working with the responseWriter
type GhastController struct{}

// Success helper for writing out a 200 status and any fmt.Fprint-able interface.
func (g GhastController) Success(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, text)
}

// NotFound helper for writing out a 404 status and any fmt.Fprint-able interface.
func (g GhastController) NotFound(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, text)
}

// BadRequest helper for writing out a 400 status and any fmt.Fprint-able interface.
func (g GhastController) BadRequest(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, text)
}

// Unauthorized helper for writing out a 401 status and any fmt.Fprint-able interface.
func (g GhastController) Unauthorized(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, text)
}

// Forbidden helper for writing out a 403 status and any fmt.Fprint-able interface.
func (g GhastController) Forbidden(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprint(w, text)
}

// InternalServerError helper for writing out a 500 status and any fmt.Fprint-able interface.
func (g GhastController) InternalServerError(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, text)
}
