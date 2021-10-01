package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "reflect"
)

// Move to internal package
func writeOut(w http.ResponseWriter, output interface{}) error {
    if reflect.ValueOf(output).Kind() != reflect.Struct {
        fmt.Fprint(w, output)
    } else {
        json, err := json.Marshal(output)
        if err != nil {
            (w).WriteHeader(http.StatusInternalServerError)
            return err
        }

        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, string(json))
    }

    return nil
}

// Success helper for writing out a 200 status and any fmt.Fprint-able interface.
func (g GhastController) Success(w http.ResponseWriter, text interface{}) {
    w.WriteHeader(http.StatusOK)
    writeOut(w, text)
}

// NotFound helper for writing out a 404 status and any fmt.Fprint-able interface.
func (g GhastController) NotFound(w http.ResponseWriter, text interface{}) {
    (w).WriteHeader(http.StatusNotFound)
    writeOut(w, text)
}

// BadRequest helper for writing out a 400 status and any fmt.Fprint-able interface.
func (g GhastController) BadRequest(w http.ResponseWriter, text interface{}) {
    w.WriteHeader(http.StatusBadRequest)
    writeOut(w, text)
}

// Unauthorized helper for writing out a 401 status and any fmt.Fprint-able interface.
func (g GhastController) Unauthorized(w http.ResponseWriter, text interface{}) {
    w.WriteHeader(http.StatusUnauthorized)
    writeOut(w, text)
}

// Forbidden helper for writing out a 403 status and any fmt.Fprint-able interface.
func (g GhastController) Forbidden(w http.ResponseWriter, text interface{}) {
    w.WriteHeader(http.StatusForbidden)
    writeOut(w, text)
}

// InternalServerError helper for writing out a 500 status and any fmt.Fprint-able interface.
func (g GhastController) InternalServerError(w http.ResponseWriter, text interface{}) {
    w.WriteHeader(http.StatusInternalServerError)
    writeOut(w, text)
}
