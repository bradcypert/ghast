package controllers

import (
	"fmt"
	"net/http"
)

// GhastController should be embedded into consumer controllers
// and provides helper functions for working with the responseWriter
type GhastController struct{}

func (g GhastController) success(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Test")
}

func (g GhastController) notFound(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Test")
}

func (g GhastController) badRequest(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Test")
}

func (g GhastController) badRequest(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Test")
}
