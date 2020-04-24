package controllers

import (
	"fmt"
	"net/http"
)

// GhastController should be embedded into consumer controllers
// and provides helper functions for working with the responseWriter
type GhastController struct{}

func (g GhastController) Success(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, text)
}

func (g GhastController) NotFound(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, text)
}

func (g GhastController) BadRequest(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, text)
}

func (g GhastController) Unauthorized(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, text)
}
